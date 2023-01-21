package main

import (
	"github/lwileczek/goBenchmarkSerialization/types"
	"log"
	"net"

	"google.golang.org/protobuf/proto"
)

// HandleProtobuf handles "Protobuf" payloads. It decodes the received Protobuf
// data into a struct, and then returns text to indicate the successful transaction.
func HandleProtobuf(conn net.Conn) {
	b := make([]byte, 1024)
	for {
		data := types.PbPayload{}
		n, err := conn.Read(b)
		if err != nil {
			log.Println("Error Reading Response from protobuf", err)
			return
		}
		err = proto.Unmarshal(b[:n], &data)
		if err != nil {
			log.Println("Protobuf: Error Unmarshalling data")
			log.Fatal(err)
			return
		}
		str := data.GetStringEntry()
		if str == "Can this be sent quickly?" {
			conn.Write([]byte("Protobuf: Successful Transfer\n"))
		}
	}
}
