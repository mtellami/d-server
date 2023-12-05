package seris

import "sync"

type Config struct {
	Port 			int
	EnableAof bool
	AofFile 	string
}

type Server struct {
	conf 			*Config
	handlers 	map[string]CommandHandler
	mu 				sync.RWMutex
}

type CommandHandler struct {
	Handler func() error
}
