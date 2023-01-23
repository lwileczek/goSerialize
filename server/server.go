package main

import (
	"io"
	"log"
	"net"
	"sync"
)

//HandleFunc - A type to handle incoming data
type HandleFunc func(net.Conn)

//Server - TCP server
type Server struct {
	ListenAddr string
	Proto      string
	Listener   net.Listener
	Handlers   map[string]HandleFunc

	// Maps are not threadsafe, so we need a mutex to control access.
	Mtx sync.RWMutex
}

//AddHandleFunc - Add a function which can be used to handle incoming requests
func (s *Server) AddHandleFunc(name string, f HandleFunc) {
	s.Mtx.Lock()
	s.Handlers[name] = f
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
	//Accept Loop: Listen Endlessly
	for {
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
		Proto:      proto,
		ListenAddr: listenAddr,
		Handlers:   make(map[string]HandleFunc),
	}
}

//RunServer - Run the server code above and do the stuff
func RunServer(addr string, protocol string) error {
	youHaveBeenServed := NewServer(addr, protocol)
	youHaveBeenServed.AddHandleFunc("GOB", HandleGob)
	youHaveBeenServed.AddHandleFunc("JSON", HandleJSON)
	youHaveBeenServed.AddHandleFunc("MSGPACK", HandleMsgPack)
	youHaveBeenServed.AddHandleFunc("PROTOBUF", HandleProtobuf)
	//Start listening.
	return youHaveBeenServed.Start()
}
