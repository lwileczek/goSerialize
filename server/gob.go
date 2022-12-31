package main

import (
	"bufio"
	"encoding/gob"
	"log"
)

// HandleGob handles the "GOB" request. It decodes the received GOB data
// into a struct, and then returns text to indicate the successful transaction.
func HandleGob(rw *bufio.ReadWriter) {
	var data Payload
	dec := gob.NewDecoder(rw)
	err := dec.Decode(&data)
	if err != nil {
		log.Println("Error decoding GOB data:", err)
		return
	}
	// Print the complex data struct to prove
	// that it correctly travelled across the wire.
	//log.Printf("Outer complexData struct: \n%#v\n", data)
	_, err = rw.WriteString("GOB: Successful Transaction.\n")
	if err != nil {
		log.Println("Cannot write to connection.\n", err)
	}
	err = rw.Flush()
	if err != nil {
		log.Println("Flush failed.", err)
	}
}
