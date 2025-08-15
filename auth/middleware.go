package auth

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func jwtParsing(tokenStr string, claims jwt.MapClaims) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtsecret, nil
	})
	return token, err
}

func VerifyJWT(c *gin.Context) {
	tokenStr := c.GetHeader("Authorization")
	tokenFromStorage := c.Query("token")
	if tokenStr == "" && tokenFromStorage == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "missing token",
		})
		c.Abort()
		return
	}
	// Bearer prefix need to manually trimmed
	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")
	tokenStr = strings.TrimSpace(tokenStr)
	claims := jwt.MapClaims{}

	token, err := jwtParsing(tokenStr, claims)

	if err != nil || !token.Valid {
		token, err = jwtParsing(tokenFromStorage, claims)
		if err != nil || !token.Valid {
			log.Printf("JWT error: %v | valid: %v", err, token.Valid)
			err_mess := fmt.Sprintf(err.Error(), !token.Valid)
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": err_mess,
			})
			c.Abort()
			return
		}
	}
	c.Set("username", claims["username"])
	c.Set("user_id", claims["user_id"])
	c.Next()

}
