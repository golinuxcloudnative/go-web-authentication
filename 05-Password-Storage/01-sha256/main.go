package main

import (
	"crypto/sha256"
	"fmt"
)

var (
	input1 = "Hello world"
	input2 = "Hello worl"
)

func main() {
	h := sha256.New()
	h.Write([]byte(input1))
	fmt.Printf("%s \t- sha256\t %x\n", input1, h.Sum(nil))
	h.Reset()
	h.Write([]byte(input2))
	fmt.Printf("%s \t- sha256\t %x\n", input2, h.Sum(nil))
}
