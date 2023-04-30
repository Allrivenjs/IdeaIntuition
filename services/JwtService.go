package services

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

var SecretKey = os.Getenv("APP_KEY")

func GenerateToken(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, err := token.SignedString([]byte(SecretKey))
	return t, err
}

func Validate(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, errors.New("invalid signing algorithm")
		}
		return []byte(SecretKey), nil
	})
}
