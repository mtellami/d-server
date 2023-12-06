package seris

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

// RESP READER
func NewResp(rd io.Reader) *Resp {
	return &Resp{reader: bufio.NewReader(rd)}
}

func (resp *Resp) readLine() (line []byte, n int, err error) {
	for {
		b, err := resp.reader.ReadByte()
		if err != nil {
			return nil, 0, err
		}
		n++;
		line = append(line, b)
		if len(line) >= 2 && line[len(line) - 2] == '\r' {
			break
		}
	}

	return line[:len(line) - 2], n, nil
}

func (resp *Resp) readInteger() (x int, n int, err error) {
	line, n, err := resp.readLine()
	if err != nil {
		return 0, 0, err
	}

	i64, err := strconv.ParseInt(string(line), 10, 64)
	if err != nil {
		return 0, n, err
	}

	return int(i64), n, nil
}

// RESP PARSER
func (resp *Resp) Read() (Value, error) {
	_t, err := resp.reader.ReadByte()
	if err != nil {
		return Value{}, err
	}

	switch _t {
	case ARRAY:
		return resp.readArray()
	case BULK:
		return resp.readBulk()
	default:
		fmt.Printf("Unknown type: %v", string(_t))
		return Value{}, nil
	}
}

// RESP ARRAY PARSER
func (resp *Resp) readArray() (Value, error) {
	v := Value{
		typ: "array",
	}

	len, _, err := resp.readInteger()
	if err != nil {
		return v, err
	}

	v.array = make([]Value, 0)
	for i := 0; i < len; i++ {
		val, err := resp.Read()
		if err != nil {
			return v, err
		}

		v.array = append(v.array, val)
	}

	return v, nil
}

// RESP BULK PARSER
func (resp *Resp) readBulk() (Value, error) {
	v := Value{
		typ: "bulk",
	}

	len, _, err := resp.readInteger()
	if err != nil {
		return v, err
	}

	bulk := make([]byte, len)

	resp.reader.Read(bulk)
	v.bulk = string(bulk)

	resp.readLine()

	return v, nil
}
