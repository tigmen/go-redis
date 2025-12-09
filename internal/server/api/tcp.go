package api

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	Addr   string
	ln     net.Listener
	quitCh chan struct{}
}

func NewServer(addr string) *Server {
	server := &Server{
		Addr:   addr,
		quitCh: make(chan struct{}),
	}

	return server
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return fmt.Errorf("Failed start listening: %w", err)
	}
	defer ln.Close()

	s.ln = ln

	go s.acceptLoop()

	<-s.quitCh

	return nil
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			log.Printf("Failed to accept: %s", err)
			continue
		}

		go s.readLoop(conn)
	}
}

func (s *Server) readLoop(conn net.Conn) {
	buf := make([]byte, 2048)
	defer conn.Close()

	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatalf("Read error: %s", err)
			continue
		}
		msg := buf[:n]

		fmt.Printf("%d: %s", n, string(msg))

		_, err = conn.Write(msg)
		if err != nil {
			log.Printf("Write error: %s", err)
			continue
		}

	}
}
