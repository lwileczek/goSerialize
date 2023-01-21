package main

import (
	"bufio"
	"encoding/json"
	"github/lwileczek/goBenchmarkSerialization/types"
	"log"
	"net"
)

// HandleJSON handles the "JSON" request. It decodes the received JSON encoded
// data into a struct, and then returns text to indicate the successful transaction.
func HandleJSON(conn net.Conn) {
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	for {
		//Will this reset the payload for me?
		var data types.Payload
		dec := json.NewDecoder(rw)
		err := dec.Decode(&data)
		if err != nil {
			log.Println("Error decoding JSON data:", err)
			_, err = rw.WriteString("JSON: Failed Transaction.\n")
			if err != nil {
				log.Println("JSON: Flush failed after error", err)
			}
			return
		}
		// Print the complex data struct to prove
		// that it correctly travelled across the wire.
		//log.Printf("Outer complexData struct: \n%#v\n", data)
		_, err = rw.WriteString("JSON: Successful Transaction.\n")
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
