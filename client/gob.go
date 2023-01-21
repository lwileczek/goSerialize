package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"github/lwileczek/goBenchmarkSerialization/types"
	"log"
)

//GOBSend - Serialize and send data through a network connection Via Go GOBs
func GOBSend(p *types.Payload, rw *bufio.ReadWriter) (bool, error) {
	enc := gob.NewEncoder(rw)
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

	//log.Println("Read the reply.")
	response, err := rw.ReadBytes('\n')
	if err != nil {
		log.Println(err, "Client: Failed to read the reply: '"+string(response)+"'")
		return false, err
	}
	//TODO: Replace with something I like more
	if bytes.Contains(response, []byte("Success")) {
		return true, nil
	}
	return false, nil
}
