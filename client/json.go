package main

import (
	"bufio"
	"encoding/json"
	"log"
	"strings"
)

//JSONSerializer - A structure to represent our JSON Serializer
type JSONSerializer struct {
	Serializer
}

//Send - Serialize and send data through a network connection
func (js *JSONSerializer) Send(p *Payload, rw *bufio.ReadWriter) {
	rw.WriteByte('j')
	enc := json.NewEncoder(rw)
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
		js.Count++
	}
}
