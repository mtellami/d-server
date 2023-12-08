package seris

import (
	"bufio"
	"os"
)

func NewAof(path string) (*Aof, error) {
	file, err := os.OpenFile(path, os.O_CREATE | os.O_RDWR | os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	// File sync latter on GoRoutine every 5 Sec
	return &Aof{file: file, rd: bufio.NewReader(file)}, nil
}

func (aof *Aof) Close() error {
	// Mutexi
	return aof.file.Close()
}

func (aof *Aof) Write(value Value) error {
	_, err := aof.file.Write(value.Marshal())
	if err != nil {
		return err
	}

	return nil
}

func (aof *Aof) Read() error {
	var values []Value

	reader := NewReader(aof.rd)
	for {
		v, err := reader.Read()
		if err != nil {
			break
		}
		values = append(values, v)
	}

	for _, value := range values {
		_, err := execute(value)
		if err != nil {
			return err
		}
	}

	return nil
}
