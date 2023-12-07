package seris

import (
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
)

// Create new Server
func NewServer(config *Config) *Server {
	handlers := make(map[string]func([]Value) Value)

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
	port := server.conf.Port
	fmt.Println("Server Listening on port:", port)

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	defer ln.Close()

	connection, err := ln.Accept()
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer connection.Close()

	for {
		err = server.handleConn(connection)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
		}
	}
	return nil
}

// Handle connections
func (server *Server) handleConn(connection net.Conn) error {
	reader := NewReader(connection)

	value, err := reader.Read()
	if err != nil {
		return err
	}

	if value.typ != "array" || len(value.array) == 0 {
		return errors.New("Invalid request, expected array")
	}

	command := strings.ToUpper(value.array[0].bulk)
	args := value.array[1:]

	writer := NewWriter(connection)

	handler, ok := server.handlers[command]
	if !ok {
		writer.Write(Value{typ: "string", str: ""})
		return errors.New(fmt.Sprintf("Invalid command: %s", command))
	}

	result := handler(args)
	writer.Write(result)

	return nil
}
