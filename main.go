package main

import (
	"chat-app/auth"
	"chat-app/controller"
	"chat-app/views"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.LoadHTMLGlob("templates/*")
	r.POST("/register", controller.Register)
	r.GET("/login", views.LoadLoginPage)
	r.POST("/login", controller.Login)
	r.GET("/chat", views.ChatPage)

	protected := r.Group("/")
	protected.Use(auth.VerifyJWT)
	protected.GET("/profile", controller.ViewProfile)
	protected.GET("/messages/:user_id", controller.GetChatHistory)
	protected.GET("/ws", controller.HttpUpgrader)
	r.Run("localhost:5000") // listen and serve on 0.0.0.0:8080
	// gin.HandlerFunc(controller.Abracadabra)
}
