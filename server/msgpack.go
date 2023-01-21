package main

import (
	"bufio"
	"github/lwileczek/goBenchmarkSerialization/types"
	"log"
	"net"

	"github.com/vmihailenco/msgpack/v5"
)

//HandleMsgPack Reads binary data from buffer and decode it using message pack
func HandleMsgPack(conn net.Conn) {
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	for {
		//Will this reset the payload for me?
		var data types.Payload
		dec := msgpack.NewDecoder(rw)
		err := dec.Decode(&data)
		if err != nil {
			log.Println("Error decoding MsgPack data:", err)
			_, err = rw.WriteString("MsgPack: Failed Transaction.\n")
			if err != nil {
				log.Println("MsgPack: Flush failed after error", err)
			}
			return
		}
		// Print the complex data struct to prove
		// that it correctly travelled across the wire.
		//log.Printf("Outer complexData struct: \n%#v\n", data)
		_, err = rw.WriteString("MsgPack: Successful Transaction.\n")
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
