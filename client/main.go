package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"
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
	encodeAndSend encodeAndSend
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

//Summarize the result of the run
//TODO: Add more summary output
func Summarize(time int, serializers []*Serializer) {
	floatTime := float64(time)
	for _, s := range serializers {
		fmt.Printf("Summary Statistics for: %s\n", s.Name)
		fmt.Printf("%d requests made in %d seconds\n", s.Count, time)
		fmt.Printf("%.2f requests per second\n", float64(s.Count)/floatTime)
	}
}

func main() {
	var runtime int
	var protocol, bindAddress string
	flag.StringVar(&bindAddress, "a", "127.0.0.1:8900", "The Bind Address")
	flag.StringVar(&protocol, "p", "tcp", "The connection protocol")
	flag.IntVar(&runtime, "t", 10, "The number of time, in seconds, the test will run for")
	flag.Parse()
	var cereals []*Serializer = []*Serializer{
		{
			Name:          "JSON",
			Protocol:      protocol,
			Addr:          bindAddress,
			encodeAndSend: JSONSend,
			Count:         0,
		},
		{
			Name:          "GOB",
			Protocol:      protocol,
			Addr:          bindAddress,
			encodeAndSend: GOBSend,
			Count:         0,
		},
	}
	stopCh := make(chan struct{})
	var wg sync.WaitGroup
	for _, serializer := range cereals {
		wg.Add(1)
		serializer := serializer // https://golang.org/doc/faq#closures_and_goroutines
		go func() {
			for {
				select {
				case _, ok := <-stopCh:
					if !ok {
						wg.Done()
						return
					}
				default:
					data := serializer.DataGen()
					rw, err := serializer.OpenConn()
					if err != nil {
						log.Fatal("Could not open a connection for JSON", err)
					}
					data.SerializationMethod = serializer.Name
					success, err := serializer.encodeAndSend(&data, rw)
					if success {
						serializer.Count++
					}
					serializer.Connection.Close()
				}
			}
		}()
	}
	// moderator
	go func(t int) {
		checkpoint := t / 5
		for w := 0; w < t; w++ {
			if w%checkpoint == 0 {
				log.Printf("(%d/%d) seconds complete, still working...\n", w, t)
			}
			time.Sleep(1 * time.Second)
		}
		log.Println("Stopping our feeders")
		close(stopCh)
	}(runtime)
	log.Println("Waiting on final request")
	wg.Wait()

	log.Println("Fin!")
	Summarize(runtime, cereals)
}

//Dynamic Select Statements: https://stackoverflow.com/questions/19992334/how-to-listen-to-n-channels-dynamic-select-statement#answer-19992525
