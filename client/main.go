package main

import (
	"bytes"
	"flag"
	"fmt"
	"github/lwileczek/goBenchmarkSerialization/types"
	"log"
	"math/rand"
	"net"
	"os"
	"sync"
	"time"

	"google.golang.org/protobuf/proto"
)

//Summarize the result of the run
//TODO: Add more summary output
func Summarize(time int, serializers []*types.Serializer) {
	floatTime := float64(time)
	for _, s := range serializers {
		fmt.Fprintf(os.Stdout, "Summary Statistics for: \t\033[0;37m%s\033[0m\n", s.Name)
		fmt.Printf("%d requests made in %d seconds\n", s.Count, time)
		fmt.Printf("%.2f requests per second\n\n", float64(s.Count)/floatTime)
	}
}

func main() {
	var runtime int
	var protocol, bindAddress, bindPort string
	flag.StringVar(&bindAddress, "a", "127.0.0.1", "The Bind Address")
	flag.StringVar(&bindPort, "port", "8900", "The Bind Address")
	flag.StringVar(&protocol, "p", "tcp", "The connection protocol")
	flag.IntVar(&runtime, "t", 10, "The number of time, in seconds, the test will run for")
	flag.Parse()
	var cereals []*types.Serializer = []*types.Serializer{
		{
			Name:          "JSON",
			Protocol:      protocol,
			Addr:          bindAddress + ":" + bindPort,
			EncodeAndSend: JSONSend,
			Count:         0,
		},
		{
			Name:          "GOB",
			Protocol:      protocol,
			Addr:          bindAddress + ":" + bindPort,
			EncodeAndSend: GOBSend,
			Count:         0,
		},
		{
			Name:          "MSGPACK",
			Protocol:      protocol,
			Addr:          bindAddress + ":" + bindPort,
			EncodeAndSend: MsgPackSend,
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
						log.Fatal("Could not open a connection", err)
					}
					data.SerializationMethod = serializer.Name
					success, err := serializer.EncodeAndSend(&data, rw)
					if success {
						serializer.Count++
					}
					serializer.Connection.Close()
				}
			}
		}()
	}
	protoCh := make(chan int)
	go func() {
		protoCount := 0
		conn, err := net.Dial(protocol, bindAddress+":"+bindPort)
		if err != nil {
			log.Println(err)
			return
		}
		defer conn.Close()
		for {
			select {
			case _, ok := <-stopCh:
				if !ok {
					protoCh <- protoCount
					return
				}
			default:
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

				data := types.PbPayload{
					StringEntry:   "Can this be sent quickly?",
					SmallInteger:  uint32(rand.Intn(256)),
					NormalInteger: int64(rand.Int()),
					Boolean:       true,
					SomeFloat:     rand.Float32(),
					IntArray:      int32Arry,
					Chart:         hashmap,
					SubShop: &types.SubStruct{
						Cat:     "Maine Coon",
						Feeling: "Joy",
					},
					SerializationMethod: "Protobuf",
				}
				//https://pkg.go.dev/google.golang.org/protobuf/proto
				byteData, err := proto.Marshal(&data)
				if err != nil {
					log.Fatal("Error Marshalling data", err)
				}
				log.Println(byteData)
				conn.Write(byteData)
				b := make([]byte, 1024)
				_, err = conn.Read(b)
				if err != nil {
					log.Println("Error Reading Response from protobuf")
				}
				if bytes.Contains(b, []byte("Success")) {
					protoCount++
				}
				//rw, err := serializer.OpenConn()
				//if err != nil {
				//	log.Fatal("Could not open a connection for JSON", err)
				//}
				//data.SerializationMethod = serializer.Name
				//success, err := serializer.encodeAndSend(&data, rw)
				//if success {
				//	serializer.Count++
				//}
				//serializer.Connection.Close()
			}
		}
	}()
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

	protoResults := <-protoCh
	fmt.Fprintf(os.Stdout, "Summary Statistics for: \t\033[0;37m%s\033[0m\n", "Protobuf")
	fmt.Printf("%d requests made in %d seconds\n", protoResults, runtime)
	fmt.Printf("%.2f requests per second\n\n", float64(protoResults)/float64(runtime))
}

//Dynamic Select Statements: https://stackoverflow.com/questions/19992334/how-to-listen-to-n-channels-dynamic-select-statement#answer-19992525
