package main

import (
	"bufio"
	"bytes"
	"github/lwileczek/goBenchmarkSerialization/types"
	"log"

	"github.com/vmihailenco/msgpack/v5"
)

//MsgPackSend - Encode using MessagePack and send data over the wire
func MsgPackSend(p *types.Payload, rw *bufio.ReadWriter) (bool, error) {
	enc := msgpack.NewEncoder(rw)
	err := enc.Encode(p)
	if err != nil {
		log.Println(err, "Encode failed for struct: %#v", *p)
		return false, err
	}
	err = rw.Flush()
	if err != nil {
		log.Println(err, "Flush failed.")
		return false, err
	}

	response, err := rw.ReadBytes('\n')
	if err != nil {
		log.Println(err, "Client: Failed to read the reply: '"+string(response)+"'")
		return false, err
	}
	//TODO: Replace with something I like more
	if bytes.Contains(response, []byte("Success")) {
		return true, nil
	}
	log.Println("Couldn't read response:\n", string(response))
	return false, nil
}
