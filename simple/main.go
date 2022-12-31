package main

import (
	"fmt"
	"time"
)

const (
	//Proto Connection protocol
	Proto = "tcp"
	//Host server
	Host = "127.0.0.1"
	//Port on host
	Port = ":8900"
)

//SubStruct A sub structure to be nested in the example
type SubStruct struct {
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

//ExampleGob An example struct that can be sent back and forth
type ExampleGob struct {
	String  string         `json:"string"`
	Number  int            `json:"number"`
	Decimal float32        `json:"decimal"`
	Chart   map[string]int `json:"chart"`
	Hero    SubStruct      `json:"hero"`
}

func main() {
	go RunServer(Host+Port, Proto)
	time.Sleep(1 * time.Second)
	go Client(Host, Port, Proto)

	fmt.Scanln()
}
