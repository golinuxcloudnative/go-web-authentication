package main

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	utils "github.com/golinuxcloudnative/go-web-authentication/99-utils"
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
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenString, err := token.SignedString([]byte(JWT_KEY))
	if err != nil {
		return "", fmt.Errorf("could not signing token, %w", err)
	}
	return tokenString, nil
}

func parseToken(signedToken string) (*UserClaims, error) {
	claims := UserClaims{}
	t, err := jwt.ParseWithClaims(signedToken, &claims, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("Invalid signing algorithm")
		}
		return []byte(JWT_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	if !t.Valid {
		return nil, err
	}

	claims = *t.Claims.(*UserClaims)

	return &claims, nil
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

	verified, err := parseToken(token)

	if err != nil {
		log.Fatalf("error to parse token, %v", err)
	}

	utils.PrettyPrint(verified)

}
