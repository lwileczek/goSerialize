package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"github/lwileczek/goBenchmarkSerialization/types"
	"log"
)

//JSONSend - Encode using JSON and send data over the wire
func JSONSend(p *types.Payload, rw *bufio.ReadWriter) (bool, error) {
	enc := json.NewEncoder(rw)
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
