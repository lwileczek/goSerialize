package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"log"
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

	//log.Println("Read the reply.")
	response, err := rw.ReadBytes('\n')
	if err != nil {
		log.Println(err, "Client: Failed to read the reply: '"+string(response)+"'")
		return
	}
	//TODO: Replace with something I like more
	if bytes.Contains(response, []byte("Success")) {
		g.Count++
	}
}

//Process incoming data and send it out to server. Keep count of requests
func (g *GOBSerializer) Process() {
	for data := range g.Queue {
		rw, conn, err := OpenConn(g.Addr, g.Protocol)
		if err != nil {
			log.Fatal("Could not open a connection for GOB", err)
		}
		data.SerializationMethod = "GOB"
		g.Send(&data, rw)
		conn.Close()
	}
}
