package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
)

// Open connects to a TCP Address.
// It returns a TCP connection armed with a timeout and wrapped into a
// buffered ReadWriter.
func Open(addr string, proto string) (*bufio.ReadWriter, error) {
	// Dial the remote process.
	// Note that the local port is chosen on the fly. If the local port
	// must be a specific one, use DialTCP() instead.
	log.Println("Dial " + addr)
	//defer conn.Close()
	conn, err := net.Dial(proto, addr)
	if err != nil {
		log.Println(err, "Dialing "+addr+" failed")
		return nil, err
	}
	log.Println("Connected to the server!!")
	return bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn)), nil
}

//Client - A local client to interact with our TCP server
func Client(ip string, port string, proto string) {
	ss := SubStruct{
		Name: "Luke Skywalker",
		Age:  21,
	}
	ex := ExampleGob{
		String:  "Hello, World!",
		Number:  rand.Int(),
		Decimal: rand.Float32(),
		Hero:    ss,
		Chart: map[string]int{
			"one":        1,
			"two":        2,
			"three":      3,
			"fourty-two": 42,
		},
	}
	for j := 0; j < 5; j++ {
		rw, err := Open(ip+port, proto)
		if err != nil {
			fmt.Println(err)
			return
		}
		//Send one byte ahead of time to indicate it will be a gob
		//(Personal choice, you don't have to do this)
		rw.WriteByte('g')
		enc := gob.NewEncoder(rw)
		err = enc.Encode(ex)
		if err != nil {
			log.Println(err, "Encode failed for struct: %#v", ex)
			continue
		}
		//Buffered Writers need to call Flush() after writing,
		//so that all data is forwarded to the underlying network connection.
		err = rw.Flush()
		if err != nil {
			log.Println(err, "Flush failed.")
			continue
		}

		log.Println("Read the reply.")
		response, err := rw.ReadString('\n')
		if err != nil {
			log.Println(err, "Client: Failed to read the reply: '"+response+"'")
			continue
		}
		log.Println("Received a response:", response)
		time.Sleep(1 * time.Second)

	}
}
