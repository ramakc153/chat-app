package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ViewProfile(c *gin.Context) {
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user id not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "profile page visited",
		"user_id": username,
	})
	return
}

func Add(a, b int) {
	fmt.Println(a + b)
	return
}
