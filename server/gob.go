package main

import (
	"bufio"
	"encoding/gob"
	"github/lwileczek/goBenchmarkSerialization/types"
	"log"
	"net"
)

// HandleGob handles the "GOB" request. It decodes the received GOB data
// into a struct, and then returns text to indicate the successful transaction.
func HandleGob(conn net.Conn) {
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	for {
		//Will this reset the payload for me?
		var data types.Payload
		dec := gob.NewDecoder(rw)
		err := dec.Decode(&data)
		if err != nil {
			log.Println("Error decoding GOB data:", err)
			_, err = rw.WriteString("GOB: Failed Transaction.\n")
			if err != nil {
				log.Println("GOB: Flush failed after error", err)
			}
			return
		}
		// Print the complex data struct to prove
		// that it correctly travelled across the wire.
		//log.Printf("Outer complexData struct: \n%#v\n", data)
		_, err = rw.WriteString("GOB: Successful Transaction.\n")
		if err != nil {
			log.Println("Cannot write to connection.\n", err)
			return
		}
		err = rw.Flush()
		if err != nil {
			log.Println("Flush failed.", err)
			return
		}
	}
}
