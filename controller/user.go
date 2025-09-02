package controller

import (
	"chat-app/auth"
	"chat-app/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var requestBody User

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		panic(err)
	}
	if len(requestBody.Password) < 6 || len(requestBody.Username) < 6 {
		c.JSON(400, gin.H{
			"message": "Username or password to short",
		})
		return
	}
	if err := database.InsertUser(requestBody.Username, requestBody.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Registered Successfully",
	})

}

func Login(c *gin.Context) {
	var requestBody User
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		panic(err)
	}
	QueriedUser, err := database.GetUser(requestBody.Username, requestBody.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}
	tokenString, err := auth.GenerateJWT(QueriedUser.Id, requestBody.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed generate token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login success",
		"token":   tokenString,
		"user_id": QueriedUser.Id,
	})
}

func GetUsers(c *gin.Context) {
	username, exist := c.Get("username")
	if !exist {
		log.Println("failed to get user_id")
	}
	users, err := database.GetAllUsers(username.(string))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "got error when getallusers",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "all users found",
		"data":    users,
	})
	return
}
