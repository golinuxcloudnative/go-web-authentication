package main

import (
	"encoding/base64"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

var (
	passwd = "DoNotHardCodeCredentials"
)

func main() {
	fmt.Println("------------------------------")
	fmt.Println("Base64 is not encryption, you can easily decode the following text. It always creates the same string.")
	fmt.Println(base64.StdEncoding.EncodeToString([]byte(passwd)))
	fmt.Println("------------------------------")
	fmt.Println("bcrypt is one of good options. Each hash of the same string is different.")
	hashedPassword, err := hashPassword(passwd)
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println(string(hashedPassword))
	fmt.Println("------------------------------")
	fmt.Println("Input the password: ")
	var passwordInput string
	fmt.Scanf("%s", &passwordInput)
	err = hashCompare(passwordInput, hashedPassword)
	if err != nil {
		fmt.Println("It's not the same password")
	} else {
		fmt.Println("It's the same password")
	}

}

func hashPassword(password string) ([]byte, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("Error while generating bcrypt hash from password: %w", err)
	}
	return bs, nil
}

func hashCompare(password string, hashedPassword []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		return fmt.Errorf("Invalid password: %w", err)
	}
	return nil
}
