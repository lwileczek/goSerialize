package main

import (
	"flag"
	"log"
	"math/rand"
	"sync"
	"time"
)

//Serializer - A serializaiton format to be tested
type Serializer struct {
	Name  string
	Queue chan Payload
}

// Should I make a more complex data type? Might be annoying with the serialization.

//Payload How we will be sending the data to the server
type Payload struct {
	StringEntry   string
	SmallInteger  uint8
	NormalInteger int
	Boolean       bool
	SomeFloat     float32
	IntArray      []int

	SerializationMethod string //MsgPack, JSON, BSON, Protobuf, etc.

}

var (
	jsonizer = Serializer{
		Name:  "JSON",
		Queue: make(chan Payload),
	}
	protobuffer = Serializer{
		Name:  "Protobuf",
		Queue: make(chan Payload),
	}
)

func main() {
	var feeders, runtime, queueSize int
	flag.IntVar(&feeders, "f", 1, "The number of feeder functions adding tasks to queue")
	flag.IntVar(&runtime, "t", 10, "The number of time, in seconds, the test will run for")
	flag.IntVar(&queueSize, "s", 1, "The feeders queue capacity")
	flag.Parse()
	var cereals []Serializer = []Serializer{jsonizer, protobuffer}
	dataCh := make(chan Payload, 4)
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
				data := Payload{
					StringEntry:   "Can this be sent quickly?",
					SmallInteger:  uint8(rand.Intn(255)),
					NormalInteger: rand.Int(),
					Boolean:       true,
					SomeFloat:     rand.Float32(),
					IntArray:      rand.Perm(256),
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
			for range serializer.Queue {
				//Simulate work
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(700)))
			}
			wg.Done()
		}()
	}
	var jsonCount, protoCount int
	//Dynamic Select Statements: https://stackoverflow.com/questions/19992334/how-to-listen-to-n-channels-dynamic-select-statement#answer-19992525
	go func() {
		for dat := range dataCh {
			select {
			case jsonizer.Queue <- dat:
				jsonCount++
			case protobuffer.Queue <- dat:
				protoCount++
			}
		}
		close(jsonizer.Queue)
		close(protobuffer.Queue)
		log.Printf("The amount of %s procecced: %d", jsonizer.Name, jsonCount)
		log.Printf("The amount of %s procecced: %d", protobuffer.Name, protoCount)
	}()
	// moderator
	go func(t int) {
		time.Sleep(time.Duration(t) * time.Second)
		close(stopCh)
	}(runtime)
	feedWg.Wait()
	close(dataCh)
	wg.Wait()
}
