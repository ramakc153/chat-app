package auth

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
)

func WsVerifyJwt(c *gin.Context) {
	arr := websocket.Subprotocols(c.Request)
	// i will keep this for debugging purpose
	// fmt.Println("this is len arr: ", len(arr))
	// fmt.Println(arr)
	if len(arr) != 2 || arr[0] != "Bearer" {
		log.Println("error missing token")
	}
	claims := jwt.MapClaims{}
	token, err := JwtParsing(arr[1], claims)
	if err != nil || !token.Valid {
		log.Printf("error parsing jwt: %v | valid: %v\n", err.Error(), token.Valid)
		c.AbortWithStatus(http.StatusUnauthorized)
		// c.Redirect(http.StatusTemporaryRedirect, "/login")
		// c.Abort()
		return
	}

	c.Set("username", claims["username"])
	c.Set("user_id", claims["user_id"])
	c.Writer.Header().Set("Sec-WebSocket-Protocol", arr[0])
	c.Next()
}
