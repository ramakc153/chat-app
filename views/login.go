package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoadLoginPage(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		c.HTML(http.StatusOK, "login.html", nil)
		return
	}
}
