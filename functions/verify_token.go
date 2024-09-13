package functions

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func VerifyToken(tokenString string) (jwt.Claims, error) {
	secretKey := []byte(os.Getenv("JWT_USER_SECRET"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	fmt.Printf("token: %v\n", token)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	fmt.Printf("token.Claims: %v\n", token.Claims)

	return token.Claims, nil
}
