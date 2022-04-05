package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type UserClaims struct {
	jwt.StandardClaims
	SessionID int64  `json:"session-id"`
	Username  string `json:"username"`
}

func (uc *UserClaims) Valid() error {
	if !uc.VerifyExpiresAt(time.Now().Unix(), true) {
		return fmt.Errorf("Token has expired")
	}

	if uc.SessionID == 0 {
		return fmt.Errorf("Invalid session id")
	}
	return nil
}

func createToken(c *UserClaims) (string, error) {
	// HS256 (HMAC with SHA-256) - a symmetric algorithm
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenString, err := token.SignedString([]byte(JWT_KEY))
	if err != nil {
		return "", fmt.Errorf("could not signing token, %w", err)
	}
	return tokenString, nil
}

var (
	JWT_KEY = "You should keep it safe"
)

func main() {

	claim := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(50 * time.Minute).Unix(),
		},
		SessionID: 1234,
		Username:  "rafael",
	}

	token, err := createToken(&claim)
	if err != nil {
		fmt.Printf("erro: %v", err)
	} else {
		fmt.Println(token)
	}

}
