package main

import (
	"log"
)

type subStruct struct {
	Cat     string `json:"cat"`
	Feeling string `json:"feeling"`
}

//Payload The endpoint is expecting to receive
type Payload struct {
	StringEntry   string          `json:"stringEntry"`
	SmallInteger  uint8           `json:"smallInteger"`
	NormalInteger int             `json:"normalInteger"`
	Boolean       bool            `json:"booleanVal"`
	SomeFloat     float32         `json:"someFloat"`
	IntArray      []int8          `json:"intArray"`
	Chart         map[string]int8 `json:"chart"`
	SubShop       subStruct       `json:"subShop"`

	SerializationMethod string `json:"serializationMethod"` //MsgPack, JSON, BSON, Protobuf, etc.

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
