package main

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var (
	input1 = "Hello world"
	input2 = "Hello worl"
)

func main() {

	for _, i := range []int{4, 4, 10, 10, 17} {
		h, duration := hashInput(input1, i)
		fmt.Printf("input: %s \tcost: %d \tduration: %v \thash: %s\n", input1, i, duration, h)
	}

	for _, i := range []int{4, 4, 10, 10, 17} {
		h, duration := hashInput(input2, i)
		fmt.Printf("input: %s \tcost: %d \tduration: %v \thash: %s\n", input2, i, duration, h)
	}
}

func hashInput(input string, cost int) ([]byte, time.Duration) {
	start := time.Now()
	hash, _ := bcrypt.GenerateFromPassword([]byte(input), cost)
	return hash, time.Since(start)
}
