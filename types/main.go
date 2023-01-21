package types

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
)

//SubStructEx A simple structure which will be nested
type SubStructEx struct {
	Cat     string `json:"cat"`
	Feeling string `json:"feeling"`
}

//Payload How we will be sending the data to the server
type Payload struct {
	StringEntry   string           `json:"stringEntry"`
	SmallInteger  uint32           `json:"smallInteger"`
	NormalInteger int              `json:"normalInteger"`
	Boolean       bool             `json:"booleanVal"`
	SomeFloat     float32          `json:"someFloat"`
	IntArray      []int32          `json:"intArray"`
	Chart         map[string]int32 `json:"chart"`
	SubShop       SubStructEx      `json:"subShop"`

	SerializationMethod string `json:"serializationMethod"` //MsgPack, JSON, BSON, Protobuf, etc.

}

type encodeAndSend func(p *Payload, rw *bufio.ReadWriter) (bool, error)

//Serializer - A serializaiton format to be tested
type Serializer struct {
	Name          string
	Flag          rune
	Protocol      string
	Addr          string
	Count         int
	Connection    net.Conn
	EncodeAndSend encodeAndSend
}

//DataGen - Generate a single payload to be serialized and transmitted
func (s *Serializer) DataGen() Payload {
	keyCount := rand.Intn(15)
	hashmap := map[string]int32{}
	for k := 0; k < keyCount; k++ {
		hashmap[fmt.Sprintf("keyNum:%d", k)] = int32(rand.Intn(256))
	}
	intArry := rand.Perm(rand.Intn(256))
	int32Arry := make([]int32, len(intArry))
	for j := 0; j < len(intArry); j++ {
		int32Arry[j] = int32(intArry[j])
	}

	data := Payload{
		StringEntry:   "Can this be sent quickly?",
		SmallInteger:  uint32(rand.Intn(256)),
		NormalInteger: rand.Int(),
		Boolean:       true,
		SomeFloat:     rand.Float32(),
		IntArray:      int32Arry,
		Chart:         hashmap,
		SubShop: SubStructEx{
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
