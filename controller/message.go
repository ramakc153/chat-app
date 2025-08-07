package controller

import (
	"chat-app/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// import "time"

// type Message struct {
// 	Id         string    `json:"id"`
// 	SenderId   string    `json:"sender_id"`
// 	ReceiverId string    `json:"receiver_id"`
// 	Content    string    `json:"content"`
// 	Timestamp  time.Time `json:"timestamp"`
// 	Status     string    `json:"status"`
// }

func GetChatHistory(c *gin.Context) {
	receiver_id := c.Param("user_id")
	user_id, ok := c.Get("user_id")
	if !ok {
		log.Println("failed to get user_id header")
	}
	messages, err := database.GetMessageByUser(user_id.(string), receiver_id)
	if err != nil {
		log.Println("error message by user", err.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	if len(messages) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "no messages found",
			"data":    []database.DBMessage{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "messages found",
		"data":    messages,
	})
	return
}
