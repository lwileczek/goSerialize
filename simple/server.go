package main

import (
	"bufio"
	"encoding/gob"
	"encoding/json"
	"io"
	"log"
	"net"
)

//HandleFunc - A type to handle incoming data
type HandleFunc func(*bufio.ReadWriter)

//Server - TCP server
type Server struct {
	listenAddr string
	Proto      string
	Ln         net.Listener
	Handler    HandleFunc
	stopCh     chan struct{}
}

//Start the TCP server
func (s *Server) Start() error {
	ln, err := net.Listen(s.Proto, s.listenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()
	s.Ln = ln

	go s.acceptLoop()

	<-s.stopCh
	return nil
}

//ReadData - read in the data from the request made
func (s *Server) readData(conn net.Conn) {
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	defer conn.Close()
	data, err := rw.ReadByte()
	switch {
	case err == io.EOF:
		log.Println("Reached EOF - close this connection.\n   ---")
		return
	case err != nil:
		log.Println("\nError reading first byte. Got: '"+string(data)+"'\n", err)
		return
	}
	switch data {
	case 'g':
		s.Handler(rw)
		return
	default:
		log.Printf("The first byte is %c", data)
	}
}

//AcceptLoop - Accept requests
func (s *Server) acceptLoop() error {
	for {
		conn, err := s.Ln.Accept()
		if err != nil {
			log.Println("Error accepting connection\n", err)
			continue
		}
		go s.readData(conn)
	}
}

// handleGob handles the "GOB" request. It decodes the received GOB data
// into a struct.
func handleGob(rw *bufio.ReadWriter) {
	log.Print("Receive GOB data:")
	var data ExampleGob
	// Create a decoder that decodes directly into a struct variable.
	dec := gob.NewDecoder(rw)
	err := dec.Decode(&data)
	if err != nil {
		log.Println("Error decoding GOB data:", err)
		return
	}
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Println("Unable to marshal object in to json")
	}
	// Print the complexData struct and the nested one, too, to prove
	// that both travelled across the wire.
	//log.Printf("Outer complexData struct: \n%#v\n", data)
	log.Println("Data Recieved:\n", string(jsonData))
	_, err = rw.WriteString("Successful transaction.\n")
	if err != nil {
		log.Println("Cannot write to connection.\n", err)
	}
	err = rw.Flush()
	if err != nil {
		log.Println("Flush failed.", err)
	}
}

//NewServer - A constructer for the server struct
func NewServer(listenAddr string, proto string) *Server {
	return &Server{
		Proto:      proto,
		listenAddr: listenAddr,
		Handler:    handleGob,
		stopCh:     make(chan struct{}),
	}
}

//RunServer - Run the server code above and do the stuff
func RunServer(addr string, protocol string) {
	dellR620 := NewServer(addr, protocol)
	log.Fatal(dellR620.Start())
}
