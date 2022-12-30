package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"sync"
)

//HandleFunc - A type to handle incoming data
type HandleFunc func(*bufio.ReadWriter)

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
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	defer conn.Close()
	data, err := rw.ReadByte()
	switch err {
	case nil:
		break
	case io.EOF:
		log.Println("Reached EOF - close this connection.\n   ---")
		return
	default:
		log.Println("\nError reading first byte. Got: '"+string(data)+"'\n", err)
		return
	}
	var encodingType string
	switch data {
	case 'g':
		encodingType = "GOB"
	case 'j':
		encodingType = "JSON"
	default:
		log.Printf("The first byte is %c which is not part of our known cases", data)
		return
	}
	s.Mtx.RLock()
	handleCommand, ok := s.Handlers[encodingType]
	s.Mtx.RUnlock()
	if ok {
		go handleCommand(rw)
	}
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
	nvidiaDGX := NewServer(addr, protocol)
	nvidiaDGX.AddHandleFunc("GOB", HandleGob)
	nvidiaDGX.AddHandleFunc("JSON", HandleJSON)

	//Start listening.
	return nvidiaDGX.Start()
}
