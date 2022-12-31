package main

import (
	"bufio"
	"encoding/json"
	"log"
)

// HandleJSON handles the "JSON" request. It decodes the received JSON encoded
// data into a struct, and then returns text to indicate the successful transaction.
func HandleJSON(rw *bufio.ReadWriter) {
	var data Payload
	dec := json.NewDecoder(rw)
	err := dec.Decode(&data)
	if err != nil {
		log.Println("Error decoding JSON data:", err)
		return
	}
	_, err = rw.WriteString("JSON: Successful Transaction.\n")
	if err != nil {
		log.Println("Cannot write to connection.\n", err)
	}
	err = rw.Flush()
	if err != nil {
		log.Println("Flush failed.", err)
	}
}
