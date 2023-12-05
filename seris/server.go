package seris

import (
	"fmt"
	"io"
	"net"
)

// Create new Server
func NewServer(config *Config) *Server {
	handlers := make(map[string]CommandHandler)

	for command, handler := range defaultHandlers {
		handlers[command] = handler
	}

	server := &Server{
		handlers: handlers,
		conf: config,
	}

	return server
}

// Listen and accept connections
func (server *Server) Listen() error {
	fmt.Println("Server Listening on port:", server.conf.Port)

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", server.conf.Port))
	if err != nil {
		return err
	}
	defer ln.Close()

	conn, err := ln.Accept()
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer conn.Close()

	for {
		err = server.handleConn(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println(err)
				break
			}
			fmt.Println(err)
		}
	}
	return nil
}

// Handle connections
func (server *Server) handleConn(conn net.Conn) error {
	// Request
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		return err
	}

	// Response
	conn.Write([]byte("+OK\r\n"))

	return nil
}
