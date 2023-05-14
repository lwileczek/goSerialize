package benchmark

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github/lwileczek/goBenchmarkSerialization/types"
	"github/lwileczek/goBenchmarkSerialization/types/fbe"
	fbeproto "github/lwileczek/goBenchmarkSerialization/types/proto"
	"log"
	"math/rand"
	"testing"

	"github.com/vmihailenco/msgpack/v5"
	"google.golang.org/protobuf/proto"
)

func dataGen() types.Payload {
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
	return data
}

//GOBMarshal how long it takes to encode data via gobs
func BenchmarkGOBMarshal(b *testing.B) {
	data := dataGen()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		enc := gob.NewEncoder(&buf)
		err := enc.Encode(data)
		if err != nil {
			panic(err)
		}
	}
}

//GOBUnmarshal Test how long to unmashral encoded data
func BenchmarkGOBUnmarshal(b *testing.B) {
	data := dataGen()
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(data)
	if err != nil {
		panic(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var obj types.Payload
		err = gob.NewDecoder(bytes.NewReader(buf.Bytes())).Decode(&obj)
		if err != nil {
			panic(err)
		}
	}
}

//JSONMarshal how long it takes to encode data via jsons
func BenchmarkJSONMarshal(b *testing.B) {
	data := dataGen()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		err := enc.Encode(data)
		if err != nil {
			panic(err)
		}
	}
}

//JSONUnmarshal Test how long to unmashral encoded data
func BenchmarkJSONUnmarshal(b *testing.B) {
	data := dataGen()
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	err := enc.Encode(data)
	if err != nil {
		panic(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var obj types.Payload
		dec := json.NewDecoder(bytes.NewReader(buf.Bytes()))
		err = dec.Decode(&obj)
		if err != nil {
			panic(err)
		}
	}
}

//MsgPackMarshal how long it takes to encode data via msgpacks
func BenchmarkMsgPackMarshal(b *testing.B) {
	data := dataGen()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		enc := msgpack.NewEncoder(&buf)
		err := enc.Encode(data)
		if err != nil {
			panic(err)
		}
	}
}

//MsgPackUnmarshal Test how long to unmashral encoded data
func BenchmarkMsgPackUnmarshal(b *testing.B) {
	data := dataGen()
	var buf bytes.Buffer
	enc := msgpack.NewEncoder(&buf)
	err := enc.Encode(data)
	if err != nil {
		panic(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var obj types.Payload
		dec := msgpack.NewDecoder(bytes.NewReader(buf.Bytes()))
		err = dec.Decode(&obj)
		if err != nil {
			panic(err)
		}
	}
}

//ProtobufMarshal how long it takes to encode data via Protobufs
func BenchmarkProtobufMarshal(b *testing.B) {
	data := dataGen()
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
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := proto.Marshal(&protoData)
		if err != nil {
			panic(err)
		}
	}
}

//ProtobufUnmarshal Test how long to unmashral encoded data
func BenchmarkProtobufUnmarshal(b *testing.B) {
	data := dataGen()
	var buf bytes.Buffer
	enc := msgpack.NewEncoder(&buf)
	err := enc.Encode(data)
	if err != nil {
		panic(err)
	}
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
	byteData, err := proto.Marshal(&protoData)
	if err != nil {
		panic(err)
	}
	b.ResetTimer()
	//https://pkg.go.dev/google.golang.org/protobuf/proto
	for i := 0; i < b.N; i++ {
		newStructObj := types.PbPayload{}
		if err != nil {
			log.Println("Error Reading Response from protobuf", err)
			return
		}
		err = proto.Unmarshal(byteData, &newStructObj)
		if err != nil {
			log.Println("Protobuf: Error Unmarshalling data")
			log.Fatal(err)
		}

	}
}

//BenchmarkFBEMarshal how long it takes to encode data via Fast Binary Encoding FBE
func BenchmarkFBEMarshal(b *testing.B) {
	data := dataGen()
	submarine := fbeproto.NewSubStructFromFieldValues("linux", "Spicy")
	emptyPayload := fbeproto.NewPbPayloadFromFieldValues(data.StringEntry, int32(data.SmallInteger), int64(data.NormalInteger), "Fast Binary Encoding", true, data.SomeFloat, data.IntArray, data.Chart, *submarine)
	buf := fbe.NewEmptyBuffer()
	writer := fbeproto.NewPbPayloadModel(buf)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Serialize the account to the FBE stream
		if _, err := writer.Serialize(emptyPayload); err != nil {
			fmt.Println("Error serializing the data", err)
			panic("serialization error")
		}
		if ok := writer.Verify(); !ok {
			panic("verify error")
		}
	}
}
