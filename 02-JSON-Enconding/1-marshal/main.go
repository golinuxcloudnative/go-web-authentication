package main

import (
	"encoding/json"
	"fmt"
	"log"

	utils "github.com/golinuxcloudnative/go-web-authentication/99-utils"
)

type Person struct {
	FirstName string
}

func main() {

	p1 := Person{
		FirstName: "Luiz Henrique",
	}

	p2 := Person{
		FirstName: "Amora",
	}

	people := []Person{p1, p2}

	bs, err := json.Marshal(people)
	if err != nil {
		log.Fatalf("cannot marshal, %v", err)
	}

	fmt.Println(string(bs))
	utils.PrettyPrint(people)
}
