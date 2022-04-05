package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	utils "github.com/golinuxcloudnative/go-web-authentication/99-utils"
	"github.com/google/uuid"
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
	if _, ok := JWT_KEYS[currentKid]; !ok {
		return "", fmt.Errorf("key does not exist")
	}
	token.Header["kid"] = currentKid
	tokenString, err := token.SignedString(JWT_KEYS[currentKid])
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

		keyID, ok := t.Header["kid"].(string)
		if !ok {
			return nil, fmt.Errorf("key in the token does not match")
		}

		return JWT_KEYS[keyID], nil
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

func generateKey() error {
	newKey := make([]byte, 64)
	_, err := io.ReadFull(rand.Reader, newKey)
	if err != nil {
		return fmt.Errorf("could not generate the new key, %w", err)
	}

	uuid, err := uuid.NewUUID()
	if err != nil {
		return fmt.Errorf("could not generate UUID, %w", err)
	}

	JWT_KEYS[uuid.String()] = newKey
	currentKid = uuid.String()
	return err
}

func printKeys() {
	log.Println("List of keys")
	for k := range JWT_KEYS {
		fmt.Printf("\tID: %s\n", k)
	}
}

var (
	// map index should use a complex string like UUID
	JWT_KEYS = map[string][]byte{}

	currentKid = ""
)

func main() {

	printKeys()
	log.Println("Generating a new key")
	err := generateKey()
	if err != nil {
		log.Println(err)
	}
	printKeys()

	claim := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(50 * time.Minute).Unix(),
		},
		SessionID: 1234,
		Username:  "rafael",
	}

	log.Println("Creating a new token.")

	token, err := createToken(&claim)
	if err != nil {
		fmt.Printf("error: %v", err)
	} else {
		fmt.Printf("A new JWT token was created - %s\n", token)
	}

	verifiedToken, err := parseToken(token)

	if err != nil {
		log.Fatalf("Token is not a valid token. Error to parse token, %v", err)
	}

	utils.PrettyPrint(verifiedToken)

	log.Println("Generating a new key")
	err = generateKey()
	if err != nil {
		log.Println(err)
	}
	printKeys()

	log.Println("Creating a new token, now it's invalid.")

	claim = UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(-50 * time.Minute).Unix(),
		},
		SessionID: 1234,
		Username:  "rafael",
	}

	token, err = createToken(&claim)
	if err != nil {
		fmt.Printf("error: %v", err)
	} else {
		fmt.Printf("A new JWT token was created - %s\n", token)
	}

	verifiedToken, err = parseToken(token)

	if err != nil {
		log.Fatalf("Token is not a valid token. Error to parse token: %v", err)
	}

	utils.PrettyPrint(verifiedToken)

}
