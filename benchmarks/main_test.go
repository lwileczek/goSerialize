package benchmark

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"testing"

	"github/lwileczek/goBenchmarkSerialization/types"
	"github/lwileczek/goBenchmarkSerialization/types/fbe"
	fbeproto "github/lwileczek/goBenchmarkSerialization/types/proto"

    "github.com/bytedance/sonic"
	"github.com/vmihailenco/msgpack/v5"
	"google.golang.org/protobuf/proto"
)

var data []types.Payload

func setup() {
	for i := 0; i < 1_000_000; i++ {
		d := dataGen()
		data = append(data, d)
	}
}

func shutdown() {}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func dataGen() types.Payload {
	keyCount := rand.Intn(15)
	hashmap := map[string]int32{}
	for k := 0; k < keyCount; k++ {
		hashmap[fmt.Sprintf("keyNum:%d", k)] = int32(rand.Intn(256))
	}
	intArry := rand.Perm(rand.Intn(256) + 1)
	int32Arry := make([]int32, len(intArry))
	for j := 0; j < len(intArry); j++ {
		int32Arry[j] = int32(intArry[j])
	}

	cats := []string{
		"Abyssinian",
		"Aegean",
		"American Bobtail",
		"American Curl",
		"American Ringtail",
		"American Shorthair",
		"American Wirehair",
		"Aphrodite Giant",
		"Arabian Mau",
		"Asian",
		"Asian Semi-longhair",
		"Australian Mist",
		"Balinese",
		"Bambino",
		"Bengal",
		"Bombay",
		"Brazilian Shorthair",
		"British Longhair",
		"British Shorthair",
		"Burmese",
		"Burmilla",
		"California Spangled",
		"Chantilly-Tiffany",
		"Chartreux",
		"Chausie",
		"Colorpoint Shorthair",
		"Cornish Rex",
		"Cymric",
		"Cyprus",
		"Devon Rex",
	}

	feelings := []string{"Joy", "Happy", "Sad", "Hopeful", "Curious"}
	r := rand.Float64()
	data := types.Payload{
		StringEntry:   fmt.Sprintf("Can this be sent quickly?%f", r),
		SmallInteger:  uint32(rand.Intn(256)),
		NormalInteger: rand.Int(),
		Boolean:       0.51 < rand.Float32(),
		SomeFloat:     rand.Float32(),
		IntArray:      int32Arry,
		Chart:         hashmap,
		SubShop: types.SubStructEx{
			Cat:     cats[rand.Intn(len(feelings))],
			Feeling: feelings[rand.Intn(len(feelings))],
		},
	}
	return data
}

func BenchmarkGOB(b *testing.B) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)
	b.Run("Marshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var d types.Payload
			switch {
			case i < len(data):
				d = data[i]
			default:
				d = dataGen()
			}

			err := enc.Encode(d)
			if err != nil {
				panic(err)
			}
		}
	})

	b.Run("Unmarshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var obj types.Payload
			if err := dec.Decode(&obj); err != nil {
				log.Println("could not decode object", err)
				panic(err)
			}
		}
	})
}

func BenchmarkJSON(b *testing.B) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	dec := json.NewDecoder(&buf)
	b.Run("Marshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var d types.Payload
			switch {
			case i < len(data):
				d = data[i]
			default:
				d = dataGen()
			}

			err := enc.Encode(d)
			if err != nil {
				panic(err)
			}
		}
	})

	b.Run("Unmarshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var obj types.Payload
			if err := dec.Decode(&obj); err != nil {
				log.Println("could not decode object", err)
				panic(err)
			}
		}
	})
}

func BenchmarkSonicJSON(b *testing.B) {
	var buf bytes.Buffer
	enc := sonic.ConfigDefault.NewEncoder(&buf)
	dec := sonic.ConfigDefault.NewDecoder(&buf)
	b.Run("Marshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var d types.Payload
			switch {
			case i < len(data):
				d = data[i]
			default:
				d = dataGen()
			}

			err := enc.Encode(d)
			if err != nil {
				panic(err)
			}
		}
	})

	b.Run("Unmarshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var obj types.Payload
			if err := dec.Decode(&obj); err != nil {
				log.Println("could not decode object", err)
				panic(err)
			}
		}
	})
}

func BenchmarkMsgPack(b *testing.B) {
	var buf bytes.Buffer
	enc := msgpack.NewEncoder(&buf)
	dec := msgpack.NewDecoder(&buf)
	b.Run("Marshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var d types.Payload
			switch {
			case i < len(data):
				d = data[i]
			default:
				d = dataGen()
			}
			err := enc.Encode(d)
			if err != nil {
				panic(err)
			}
		}
	})

	b.Run("Unmarshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var obj types.Payload
			if err := dec.Decode(&obj); err != nil {
				log.Println("could not decode object", err)
				panic(err)
			}
		}
	})
}

func BenchmarkProtobuf(b *testing.B) {
	paylaods := make([][]byte, 0, 600_000)
	b.Run("Marshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var d types.Payload
			switch {
			case i < len(data):
				d = data[i]
			default:
				d = dataGen()
			}
			protoData := types.PbPayload{
				StringEntry:   d.StringEntry,
				SmallInteger:  uint32(d.SmallInteger),
				NormalInteger: int64(d.NormalInteger),
				Boolean:       d.Boolean,
				SomeFloat:     d.SomeFloat,
				IntArray:      d.IntArray,
				Chart:         d.Chart,
				SubShop: &types.SubStruct{
					Cat:     d.SubShop.Cat,
					Feeling: d.SubShop.Feeling,
				},
				SerializationMethod: "Protobuf",
			}
			bitz, err := proto.Marshal(&protoData)
			if err != nil {
				panic(err)
			}
			b.StopTimer()
			paylaods = append(paylaods, bitz)
			b.StartTimer()
		}
	})

	b.Run("Unmarshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var d []byte
			switch {
			case i < len(paylaods):
				d = paylaods[i]
			default:
				b.StopTimer()
				g := dataGen()
				protoData := types.PbPayload{
					StringEntry:   g.StringEntry,
					SmallInteger:  uint32(g.SmallInteger),
					NormalInteger: int64(g.NormalInteger),
					Boolean:       g.Boolean,
					SomeFloat:     g.SomeFloat,
					IntArray:      g.IntArray,
					Chart:         g.Chart,
					SubShop: &types.SubStruct{
						Cat:     "Maine Coon",
						Feeling: "Joy",
					},
					SerializationMethod: "Protobuf",
				}
				foo, err := proto.Marshal(&protoData)
				if err != nil {
					panic(err)
				}
				b.StartTimer()
				d = foo
			}

			newStructObj := types.PbPayload{}
			err := proto.Unmarshal(d, &newStructObj)
			if err != nil {
				log.Println("Protobuf: Error Unmarshalling data")
				log.Fatal(err)
			}
		}
	})
}

// BenchmarkFBEMarshal how long it takes to encode data via Fast Binary Encoding FBE
func BenchmarkFBE(b *testing.B) {
	buf := fbe.NewEmptyBuffer()
	writer := fbeproto.NewPbPayloadModel(buf)
	reader := fbeproto.NewPbPayloadModel(writer.Buffer())

	b.Run("marshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			data := dataGen()
			submarine := fbeproto.NewSubStructFromFieldValues("linux", "Spicy")
			emptyPayload := fbeproto.NewPbPayloadFromFieldValues(data.StringEntry, int32(data.SmallInteger), int64(data.NormalInteger), "Fast Binary Encoding", true, data.SomeFloat, data.IntArray, data.Chart, *submarine)
			// Serialize the account to the FBE stream
			if _, err := writer.Serialize(emptyPayload); err != nil {
				fmt.Println("Error serializing the data", err)
				panic("serialization error")
			}
			if ok := writer.Verify(); !ok {
				panic("verify fbe writer error")
			}
		}
	})

	for i := 0; i < 50_000; i++ {
		data := dataGen()
		submarine := fbeproto.NewSubStructFromFieldValues("linux", "Spicy")
		emptyPayload := fbeproto.NewPbPayloadFromFieldValues(data.StringEntry, int32(data.SmallInteger), int64(data.NormalInteger), "Fast Binary Encoding", true, data.SomeFloat, data.IntArray, data.Chart, *submarine)
		// Serialize the account to the FBE stream
		if _, err := writer.Serialize(emptyPayload); err != nil {
			fmt.Println("Error serializing the data", err)
			panic("serialization error")
		}
		if ok := writer.Verify(); !ok {
			panic("verify fbe writer error")
		}
	}

	b.Run("Unmarshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if ok := reader.Verify(); !ok {
				log.Println("The reader failed verification")
				panic("verify reader error")
			}
			if _, _, err := reader.Deserialize(); err != nil {
				log.Println(err)
				panic("unmarshal fbe error")
			}
		}
	})
}
