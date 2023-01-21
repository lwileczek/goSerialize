package files

//package main

import (
	"bufio"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github/lwileczek/goBenchmarkSerialization/types"
	"math/rand"
	"os"

	"github.com/vmihailenco/msgpack/v5"
	"google.golang.org/protobuf/proto"
)

func main() {
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

	data := types.Payload{
		StringEntry:   "Can this be sent quickly?",
		SmallInteger:  uint32(rand.Intn(256)),
		NormalInteger: rand.Int(),
		Boolean:       true,
		SomeFloat:     rand.Float32(),
		IntArray:      int32Arry,
		Chart:         hashmap,
		SubShop: types.SubStructEx{
			Cat:     "Maine Coon",
			Feeling: "Joy",
		},
	}
	writeGOB(&data)
	writeJSON(&data)
	writeMsgPack(&data)
	writeProtobuf(&data)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func writeGOB(data *types.Payload) {
	f, err := os.Create("./data.gob.bin")
	check(err)
	defer f.Close()
	w := bufio.NewWriter(f)
	enc := gob.NewEncoder(w)
	err = enc.Encode(data)
	check(err)
	w.Flush()
}

func writeJSON(data *types.Payload) {
	f, err := os.Create("./data.json")
	check(err)
	defer f.Close()
	w := bufio.NewWriter(f)
	enc := json.NewEncoder(w)
	err = enc.Encode(data)
	check(err)
	w.Flush()
}

func writeMsgPack(data *types.Payload) {
	f, err := os.Create("./data.msgpack.bin")
	check(err)
	defer f.Close()
	w := bufio.NewWriter(f)
	enc := msgpack.NewEncoder(w)
	err = enc.Encode(data)
	check(err)
	w.Flush()
}

func writeProtobuf(data *types.Payload) {
	f, err := os.Create("./data.protobuf.bin")
	check(err)
	defer f.Close()

	protoData := types.PbPayload{
		StringEntry:   "Can this be sent quickly?",
		SmallInteger:  uint32(data.SmallInteger),
		NormalInteger: int64(data.NormalInteger),
		Boolean:       true,
		SomeFloat:     data.SomeFloat,
		IntArray:      data.IntArray,
		Chart:         data.Chart,
		SubShop: &types.SubStruct{
			Cat:     "Maine Coon",
			Feeling: "Joy",
		},
		SerializationMethod: "Protobuf",
	}
	//https://pkg.go.dev/google.golang.org/protobuf/proto
	byteData, err := proto.Marshal(&protoData)
	check(err)
	byteCount, err := f.Write(byteData)
	check(err)
	fmt.Printf("wrote %d bytes in protobuf\n", byteCount)
}
