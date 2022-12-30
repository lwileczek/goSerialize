package main

import (
	"bufio"
	"encoding/gob"
	"log"
	"strings"
)

//GOBSerializer - A structure to represent our JSON Serializer
type GOBSerializer struct {
	Serializer
}

//Send - Serialize and send data through a network connection
func (g *GOBSerializer) Send(p *Payload, rw *bufio.ReadWriter) {
	rw.WriteByte('g')
	enc := gob.NewEncoder(rw)
	err := enc.Encode(p)
	if err != nil {
		log.Println(err, "Encode failed for struct: %#v", *p)
		return
	}
	err = rw.Flush()
	if err != nil {
		log.Println(err, "Flush failed.")
		return
	}

	log.Println("Read the reply.")
	response, err := rw.ReadString('\n')
	if err != nil {
		log.Println(err, "Client: Failed to read the reply: '"+response+"'")
		return
	}
	if strings.Contains(response, "Success") {
		g.Count++
	}
}
