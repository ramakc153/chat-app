package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ChatPage(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		c.HTML(http.StatusOK, "chat.html", nil)
		return
	}
}
