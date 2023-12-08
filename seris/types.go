package seris

import (
	"bufio"
	"io"
	"os"
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

type Aof struct {
	file 	*os.File
	rd 		*bufio.Reader
	mu  	sync.Mutex
}

type Server struct {
	conf 			*Config
	aof 			*Aof
}

type Data struct {
	mu 				sync.RWMutex
	hmu 			sync.RWMutex
	SETs 			map[string]string
	HSETs 		map[string]map[string]string
}

type Reader struct {
	reader *bufio.Reader
}

type Writer struct {
	writer io.Writer
}

type Value struct {
	typ	 	string
	str 	string
	num 	int
	bulk 	string
	array []Value
}
