package main

import (
	"log"
)

//Payload - the expected data structure from the client
type Payload struct {
	StringEntry   string
	SmallInteger  uint8
	NormalInteger int
	Boolean       bool
	SomeFloat     float32
	IntArray      []int
	Chart         map[string]int8

	SerializationMethod string //MsgPack, JSON, BSON, Protobuf, etc.

}

func main() {
	cfg, err := Configure()
	if err != nil {
		log.Fatal("Error configuring the server:", err)
	}
	log.Println("Constructing a server with the following configuration:", cfg)
	err = RunServer(cfg.BindAddr, cfg.Protocol)
	if err != nil {
		log.Fatal("Error running the server:", err)
	}
}
