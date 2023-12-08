package seris

import (
	"bufio"
	"os"
	"time"
)

func NewAof(path string) (*Aof, error) {
	file, err := os.OpenFile(path, os.O_CREATE | os.O_RDWR | os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	aof := &Aof{
		file: file,
		rd: 	bufio.NewReader(file),
	}

	// GoRoutine to Sync AOF buffer to disk
	go func() {
		for {
			aof.mu.Lock()
			aof.file.Sync()
			aof.mu.Unlock()
			time.Sleep(time.Second * 5)
		}
	}()

	return aof, nil
}

func (aof *Aof) Close() error {
	aof.mu.Lock()
	defer aof.mu.Unlock()

	return aof.file.Close()
}

func (aof *Aof) Write(value Value) error {
	aof.mu.Lock()
	defer aof.mu.Unlock()

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
