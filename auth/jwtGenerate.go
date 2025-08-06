package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtsecret = []byte("secret-key")

func GenerateJWT(id, username string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  id,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtsecret)
}
