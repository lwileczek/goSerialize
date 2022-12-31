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

//Payload How we will be sending the data to the server
type Payload struct {
	StringEntry   string
	SmallInteger  uint8
	NormalInteger int
	Boolean       bool
	SomeFloat     float32
	IntArray      []int
	Chart         map[string]int8

	SerializationMethod string //MsgPack, JSON, BSON, Protobuf, etc.

}

//Sender An interface built on top of Serializers
//which indicates it'll also send data across the wire
type Sender interface {
	Send(p *Payload, rw *bufio.ReadWriter)
	Process()
}

//Serializer - A serializaiton format to be tested
type Serializer struct {
	Name     string
	Queue    chan Payload
	Protocol string
	Addr     string
	Count    int
}

//TODO: Dynamically create and add Serializers
//Lookup reflector for switch cases
var (
	jsonizer = &JSONSerializer{
		Serializer{
			Name:     "JSON",
			Queue:    make(chan Payload),
			Protocol: "tcp",
			Addr:     "127.0.0.1:8000",
		},
	}
	gobber = &GOBSerializer{
		Serializer{
			Name:     "JSON",
			Queue:    make(chan Payload),
			Protocol: "tcp",
			Addr:     "127.0.0.1:8000",
		},
	}
)

// OpenConn connects to a PROTO [tcp] Address.
// It returns a connection armed with a timeout and wrapped into a
// buffered ReadWriter.
func OpenConn(addr string, proto string) (*bufio.ReadWriter, net.Conn, error) {
	// Dial the remote process.
	// Note that the local port is chosen on the fly. If the local port
	// must be a specific one, use DialTCP() instead.
	//log.Println("Dial " + addr)
	//defer conn.Close()
	conn, err := net.Dial(proto, addr)
	if err != nil {
		log.Println(err, " Dialing "+addr+" failed")
		return nil, nil, err
	}
	//log.Println("Connected to the server!!")
	return bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn)), conn, nil
}

//Summarize the result of the run
func Summarize(time int) {
	jCount := jsonizer.Count
	gCount := gobber.Count
	totalCount := jCount + gCount
	floatCount := float64(totalCount)
	floatTime := float64(time)
	jPerc := float64(jCount) / floatCount
	gPerc := float64(gCount) / floatCount
	jRPS := float64(jCount) / floatTime
	gRPS := float64(gCount) / floatTime

	fmt.Printf("%d total requests made in %d seconds\n", totalCount, time)
	fmt.Printf("%.2f requests per second\n", floatCount/floatTime)
	fmt.Printf("%.2f percent of the requests were JSON and %.2f were Gob\n", jPerc, gPerc)
	fmt.Printf("Requests per second - JSON: %.2f | GOB: %.2f\n", jRPS, gRPS)
}

func main() {
	var feeders, runtime, queueSize int
	flag.IntVar(&feeders, "f", 1, "The number of feeder functions adding tasks to queue")
	flag.IntVar(&runtime, "t", 10, "The number of time, in seconds, the test will run for")
	flag.IntVar(&queueSize, "s", 1, "The feeders queue capacity")
	flag.Parse()
	var cereals []Sender = []Sender{jsonizer, gobber}
	dataCh := make(chan Payload, 3)
	stopCh := make(chan bool)
	var feedWg sync.WaitGroup
	for fd := 0; fd < feeders; fd++ {
		feedWg.Add(1)
		go func() {
			for {
				select {
				case <-stopCh:
					feedWg.Done()
					return
				default:
				}
				keyCount := rand.Intn(15)
				hashmap := map[string]int8{}
				for k := 0; k < keyCount; k++ {
					hashmap[fmt.Sprintf("keyNum:%d", k)] = int8(rand.Intn(256))
				}
				data := Payload{
					StringEntry:   "Can this be sent quickly?",
					SmallInteger:  uint8(rand.Intn(256)),
					NormalInteger: rand.Int(),
					Boolean:       true,
					SomeFloat:     rand.Float32(),
					IntArray:      rand.Perm(256),
					Chart:         hashmap,
				}
				dataCh <- data
			}
		}()
	}
	var wg sync.WaitGroup
	for _, serializer := range cereals {
		wg.Add(1)
		serializer := serializer // https://golang.org/doc/faq#closures_and_goroutines
		go func() {
			serializer.Process()
			wg.Done()
		}()
	}
	//Dynamic Select Statements: https://stackoverflow.com/questions/19992334/how-to-listen-to-n-channels-dynamic-select-statement#answer-19992525
	go func() {
		for dat := range dataCh {
			select {
			case jsonizer.Queue <- dat:
			case gobber.Queue <- dat:
			}
		}
		close(jsonizer.Queue)
		close(gobber.Queue)
		//Make dynamic
		//for _, serializer := range cereals {
		//	close(serializer.Queue)
		//}
	}()
	// moderator
	go func(t int) {
		checkpoint := t / 5
		for w := 0; w < t; w++ {
			if w%checkpoint == 0 {
				log.Printf("(%d/%d) second complete, still working...\n", w, t)
			}
			time.Sleep(1 * time.Second)
		}
		close(stopCh)
		log.Println("Stopping our feeders")
	}(runtime)
	feedWg.Wait()
	close(dataCh)
	log.Println("Waiting on final request")
	wg.Wait()

	log.Println("Fin!")
	Summarize(runtime)
}
