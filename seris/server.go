package seris

import (
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
)

// Create new Server
func NewServer(config *Config) (*Server, error) {

	if config.EnableAof {
		aof, err := NewAof(config.AofFile)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		r := aof.Read()
		if r != nil {
			fmt.Println(r)
			return nil, r
		}
		return &Server{conf: config, aof: aof}, nil
	}

	return &Server{conf: config}, nil
}

// Listen and accept connections
func (server *Server) Listen() error {
	port := server.conf.Port
	fmt.Println("Server Listening on port:", port)

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer ln.Close()

	if server.conf.EnableAof {
		defer server.aof.Close()
	}

	for {
		connection, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			return err
		}

		// GoRoutine for every client
		go response(server, connection)
	}
}

// Client Thread
func response(server *Server, connection net.Conn) {
	defer connection.Close()

	for {
		err := server.serv(connection)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
		}
	}
}

// Handle connections
func (server *Server) serv(connection net.Conn) error {

	// REQUEST
	reader := NewReader(connection)
	value, err := reader.Read()
	if err != nil {
		return err
	}

	if value.typ != "array" || len(value.array) == 0 {
		return errors.New("Invalid request, expected array")
	}

	// EXECUTE 
	command := strings.ToUpper(value.array[0].bulk)
	response, err := execute(value)
	if err != nil {
		fmt.Println(err)
	}

	// RESPONSE
	writer := NewWriter(connection)
	writer.Write(response)

	// SAVE
	if server.conf.EnableAof && command == "SET" || command == "HSET" {
		err := server.aof.Write(value)
		if err != nil {
			return err
		}
	}

	return nil
}
