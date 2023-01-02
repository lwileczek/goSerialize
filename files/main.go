package main

import (
	"fmt"
	"math/rand"
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

func main() {
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
	fmt.Println(data)

}
