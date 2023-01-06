package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	EncodeJson()
	fmt.Println("----")
	DecodeJson()
}

func DecodeJson() {
	jsonBytes, _ := json.Marshal(Member{"Tim", 1, true})

	var mem Member
	err := json.Unmarshal(jsonBytes, &mem)
	if err != nil {
		panic(err)
	}

	fmt.Println(mem.Name, mem.Age, mem.Active)
}

type Member struct {
	Name   string
	Age    int
	Active bool
}

func EncodeJson() {
	mem := Member{"Alex", 10, true}

	jsonBytes, err := json.Marshal(mem)
	if err != nil {
		panic(err)
	}

	jsonString := string(jsonBytes)
	fmt.Println(jsonString)
}
