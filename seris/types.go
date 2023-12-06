package seris

import (
	"bufio"
	"sync"
)

const (
	STRING 	= '+'
	ERROR		= '-'
	INTEGER = ':'
	BULK 		= '$'
	ARRAY 	= '*'
)

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

type Resp struct {
	reader *bufio.Reader
}

type Value struct {
	typ	 	string
	str 	string
	num 	int
	bulk 	string
	array []Value
}
