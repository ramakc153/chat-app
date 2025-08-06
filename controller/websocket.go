package controller

import (
	"chat-app/database"
	"chat-app/utils"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all connections (use a stricter check in production)
	},
}

var connections = make(map[string]*websocket.Conn)

func HttpUpgrader(c *gin.Context) {

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Http upgrader error: ", err.Error())
		return
	}
	defer conn.Close()
	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read message error: ", err.Error())

			break
		}
		log.Println("received message: ", string(message))

		var parsed_message database.DBMessage
		err = json.Unmarshal(message, &parsed_message)
		if err != nil {
			log.Println("error when unmarshalling the message:", err.Error())
			break
		}
		// parsed_message.Id = uuid.New().String()
		parsed_message.Timestamp = utils.Generate_time()
		parsed_message.Status = "Pending"
		user_id, exists := c.Get("user_id")
		if !exists {
			log.Println("failed to retrieve user_id")
			break
		}
		connections[user_id.(string)] = conn
		receiver := parsed_message.ReceiverId
		err = database.InsertMessage(parsed_message)
		if err != nil {
			log.Println("error inserting message:", err.Error())
			break
		}
		receiverConn, exist := connections[receiver]
		if exist {
			
			err = receiverConn.WriteMessage(mt, message)
			if err != nil {
				log.Println("read message error: ", err.Error())
				break
			}
		}
		// 

	}
}

func HandleWebSocketConnection(conn *websocket.Conn)
