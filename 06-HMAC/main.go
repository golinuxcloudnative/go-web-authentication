package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"fmt"
)

var key = []byte{}

func main() {
	for i := 1; i <= 64; i++ {
		key = append(key, byte(i))
	}

	msg := "This is the message"
	bs, _ := signingMessage([]byte(msg))
	fmt.Println(string(bs))

	ok, _ := checkSignature([]byte(msg), bs)
	if ok {
		fmt.Println("Signature OK")
	} else {
		fmt.Println("Signature NOK")
	}

}

func signingMessage(message []byte) ([]byte, error) {
	//https://pkg.go.dev/crypto/sha512
	//sha512  Size = 64
	hash := hmac.New(sha512.New, key)
	_, err := hash.Write(message)
	if err != nil {
		return nil, fmt.Errorf("cannot sign the message, %w", err)
	}

	signature := hash.Sum(nil)

	return signature, nil
}

func checkSignature(msg, sig []byte) (bool, error) {
	newSig, err := signingMessage(msg)
	if err != nil {
		return false, fmt.Errorf("cannot check the signature of the message, %w", err)
	}

	same := hmac.Equal(sig, newSig)

	return same, nil
}
