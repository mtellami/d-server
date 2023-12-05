package seris

import "fmt"

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

func (server *Server) Listen() {
	fmt.Println("Server Listening on port:", server.conf.Port)
}
