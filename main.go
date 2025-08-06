package main

import (
	"chat-app/auth"
	"chat-app/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)

	protected := r.Group("/")
	protected.Use(auth.VerifyJWT)
	protected.GET("/profile", controller.ViewProfile)
	protected.GET("/ws", controller.HttpUpgrader)
	r.Run("localhost:5000") // listen and serve on 0.0.0.0:8080
	// gin.HandlerFunc(controller.Abracadabra)
}
