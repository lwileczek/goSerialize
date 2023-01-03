package types

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
)

// Should I make a more complex data type? Might be annoying with the serialization.
type subStruct struct {
	Cat     string `json:"cat"`
	Feeling string `json:"feeling"`
}

//Payload How we will be sending the data to the server
type Payload struct {
	StringEntry   string          `json:"stringEntry"`
	SmallInteger  uint8           `json:"smallInteger"`
	NormalInteger int             `json:"normalInteger"`
	Boolean       bool            `json:"booleanVal"`
	SomeFloat     float32         `json:"someFloat"`
	IntArray      []int8          `json:"intArray"`
	Chart         map[string]int8 `json:"chart"`
	SubShop       subStruct       `json:"subShop"`

	SerializationMethod string `json:"serializationMethod"` //MsgPack, JSON, BSON, Protobuf, etc.

}

type encodeAndSend func(p *Payload, rw *bufio.ReadWriter) (bool, error)

//Serializer - A serializaiton format to be tested
type Serializer struct {
	Name          string
	Protocol      string
	Addr          string
	Count         int
	Connection    net.Conn
	EncodeAndSend encodeAndSend
}

//DataGen - Generate a single payload to be serialized and transmitted
func (s *Serializer) DataGen() Payload {
	keyCount := rand.Intn(15)
	hashmap := map[string]int8{}
	for k := 0; k < keyCount; k++ {
		hashmap[fmt.Sprintf("keyNum:%d", k)] = int8(rand.Intn(256))
	}
	intArry := rand.Perm(rand.Intn(256))
	int8Arry := make([]int8, len(intArry))
	for j := 0; j < len(intArry); j++ {
		int8Arry[j] = int8(intArry[j])
	}

	data := Payload{
		StringEntry:   "Can this be sent quickly?",
		SmallInteger:  uint8(rand.Intn(256)),
		NormalInteger: rand.Int(),
		Boolean:       true,
		SomeFloat:     rand.Float32(),
		IntArray:      int8Arry,
		Chart:         hashmap,
		SubShop: subStruct{
			Cat:     "Maine Coon",
			Feeling: "Joy",
		},
	}
	return data
}

// OpenConn connects to a PROTO [tcp] Address.
// It returns a connection armed with a timeout and wrapped into a
// buffered ReadWriter.
//TODO: probably don't need to return the RW or Conn but attatch to instance
func (s *Serializer) OpenConn() (*bufio.ReadWriter, error) {
	conn, err := net.Dial(s.Protocol, s.Addr)
	if err != nil {
		log.Println(err, " Dialing "+s.Addr+" failed")
		return nil, err
	}
	//log.Println("Connected to the server!!")
	s.Connection = conn
	return bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn)), nil
}
