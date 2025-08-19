package controller

import (
	"chat-app/database"
	"chat-app/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Connections struct {
	sync.Mutex
	M map[string]*websocket.Conn
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all connections (use a stricter check in production)
	},
}

// var connections = make(map[string]*websocket.Conn)
var connections Connections

func HttpUpgrader(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Http upgrader error: ", err.Error())
		return
	}
	defer conn.Close()
	// load all pending messages
	user_id, exists := c.Get("user_id")
	if !exists {
		log.Println("failed to retrieve user_id")
	}
	connections.Lock()
	connections.M[user_id.(string)] = conn
	connections.Unlock()
	// connections[user_id.(string)] = conn
	messages, err := database.GetPendingMessage(user_id.(string))

	if err != nil {
		log.Println(err.Error())
	}
	for _, message := range messages {
		jsonMessage, err := json.Marshal(message)
		if err != nil {
			log.Println("error when marshalling message", err)
		}
		err = conn.WriteMessage(websocket.TextMessage, jsonMessage)
		if err != nil {
			log.Println("error when writing message: ", err.Error())
			return
		}
		err = database.UpdateMessageStatus(message.Id)
		if err != nil {
			log.Println(err.Error())
		}
	}
	for {
		// wait for the incoming messages
		mt, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("read message error: ", err.Error())
			}
			connections.Lock()
			delete(connections.M, user_id.(string))
			connections.Unlock()
			fmt.Println("user disconnected")

			break
		}
		log.Println("received message: ", string(message))

		var parsed_message database.DBMessage
		err = json.Unmarshal(message, &parsed_message)
		if err != nil {
			log.Println("error when unmarshalling the message:", err.Error())
			break
		}
		parsed_message.Id = uuid.New().String()
		parsed_message.SenderId = user_id.(string)
		parsed_message.Timestamp = utils.Generate_time()
		parsed_message.Status = "pending"
		receiver := parsed_message.ReceiverId
		jsonToSend, err := json.Marshal(parsed_message)
		// fmt.Println("this is parsed message", parsed_message)
		if err != nil {
			log.Println("error when marshalling parsed_message", err.Error())
		}
		err = database.InsertMessage(parsed_message)
		if err != nil {
			log.Println("error inserting message:", err.Error())
			break
		}
		// send message to receiver
		receiverConn, exist := connections.M[receiver]
		if exist {

			err = receiverConn.WriteMessage(mt, jsonToSend)
			if err != nil {
				log.Println("read message error: ", err.Error())
				break
			}
		}
		//

	}
}

// func HandleWebSocketConnection(conn *websocket.Conn){}
