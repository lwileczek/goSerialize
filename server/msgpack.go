package main

import (
	"bufio"
	"github/lwileczek/goBenchmarkSerialization/types"
	"log"

	"github.com/vmihailenco/msgpack/v5"
)

//HandleMsgPack Reads binary data from buffer and decode it using message pack
func HandleMsgPack(rw *bufio.ReadWriter) {
	var data types.Payload
	dec := msgpack.NewDecoder(rw)
	err := dec.Decode(&data)
	if err != nil {
		log.Println("Error decoding Msg Pack data:", err)
		return
	}
	_, err = rw.WriteString("MsgPack: Successful Transaction.\n")
	if err != nil {
		log.Println("Cannot write to connection.\n", err)
	}
	err = rw.Flush()
	if err != nil {
		log.Println("Flush failed.", err)
	}
}
