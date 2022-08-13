package token

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(username string) (string, error) {
	var signingKey = []byte(os.Getenv("JWT_SECRET"))

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	return token.SignedString(signingKey)
}
