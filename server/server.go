package main

import (
	"io"
	"log"
	"net"
	"sync"
	"time"
)

//HandleFunc - A type to handle incoming data
type HandleFunc func(net.Conn)

//Server - TCP server
type Server struct {
	ListenAddr    string
	Proto         string
	Listener      net.Listener
	Handlers      map[string]HandleFunc
	HandlerCounts map[string]*int

	// Maps are not threadsafe, so we need a mutex to control access.
	Mtx sync.RWMutex
}

//AddHandleFunc - Add a function which can be used to handle incoming requests
func (s *Server) AddHandleFunc(name string, f HandleFunc) {
	z := 0
	s.Mtx.Lock()
	s.Handlers[name] = f
	s.HandlerCounts[name] = &z
	s.Mtx.Unlock()
}

//Start the TCP server
func (s *Server) Start() error {
	ln, err := net.Listen(s.Proto, s.ListenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()
	s.Listener = ln
	//Check if requests are coming in
	//TODO: Remove or add cli options around this
	go func() {
		var previous int
		for {
			requestCount := 0
			time.Sleep(30 * time.Second)
			s.Mtx.Lock()
			for _, v := range s.HandlerCounts {
				requestCount += *v
			}
			s.Mtx.Unlock()
			if previous != requestCount {
				s.Mtx.Lock()
				for k, v := range s.HandlerCounts {
					log.Printf("%s : %d", k, *v)
				}
				s.Mtx.Unlock()
				log.Printf("\n\n")
				previous = requestCount
			}
		}
	}()
	//Accept Loop
	for {
		//Listen Endlessly
		conn, err := s.Listener.Accept()
		if err != nil {
			log.Println("Error accepting connection\n", err)
			continue
		}
		go s.handleRequest(conn)
	}
}

//ReadData - read in the data from the request made
func (s *Server) handleRequest(conn net.Conn) {
	data := make([]byte, 1)
	defer conn.Close()
	_, err := conn.Read(data)
	switch err {
	case nil:
		break
	case io.EOF:
		log.Println("Reached EOF - close this connection.\n   ---")
		return
	default:
		log.Printf("\nError reading first byte. Got: %s\n%s", data, err)
		return
	}
	var encodingType string
	switch rune(data[0]) {
	case 'g':
		encodingType = "GOB"
	case 'j':
		encodingType = "JSON"
	case 'm':
		encodingType = "MSGPACK"
	case 'p':
		encodingType = "PROTOBUF"
	default:
		log.Printf("The first byte is %c which is not part of our known cases", data)
		return
	}
	s.Mtx.RLock()
	handleCommand, ok := s.Handlers[encodingType]
	s.Mtx.RUnlock()
	if !ok {
		log.Printf("Could not find handler for %s", encodingType)
		return
	}
	handleCommand(conn)
}

//NewServer - A constructer for the server struct
func NewServer(listenAddr string, proto string) *Server {
	return &Server{
		Proto:         proto,
		ListenAddr:    listenAddr,
		Handlers:      make(map[string]HandleFunc),
		HandlerCounts: make(map[string]*int),
	}
}

//RunServer - Run the server code above and do the stuff
func RunServer(addr string, protocol string) error {
	nvidiaDGX := NewServer(addr, protocol)
	nvidiaDGX.AddHandleFunc("GOB", HandleGob)
	nvidiaDGX.AddHandleFunc("JSON", HandleJSON)
	nvidiaDGX.AddHandleFunc("MSGPACK", HandleMsgPack)
	nvidiaDGX.AddHandleFunc("PROTOBUF", HandleProtobuf)
	//Start listening.
	return nvidiaDGX.Start()
}
