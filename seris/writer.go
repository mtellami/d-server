package seris

import (
	"io"
	"strconv"
)

// WRITER
func NewWriter(w io.Writer) *Writer {
	return &Writer{writer: w}
}

func (writer *Writer) Write(value Value) error {
	var bytes = value.Marshal()

	_, err := writer.writer.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}

func (value Value) Marshal() []byte {
	switch value.typ {
	case "array":
		return value.marshalArray()
	case "bulk":
		return value.marshalBulk()
	case "string":
		return value.marshalString()
	case "integer":
		return value.marshalInteger()
	case "null":
		return value.marshalNull()
	case "error":
		return value.marshalError()
	default:
		return []byte{}
	}
}

// WRITER MARSHALS
func (value Value) marshalArray() []byte {
	var bytes []byte
	len := len(value.array)

	bytes = append(bytes, ARRAY)
	bytes = append(bytes, strconv.Itoa(len)...)
	bytes = append(bytes, '\r', '\n')

	for i := 0; i < len; i++ {
		bytes = append(bytes, value.array[i].Marshal()...)
	}

	return bytes
}

func (value Value) marshalBulk() []byte {
	var bytes []byte
	bytes = append(bytes, BULK)
	bytes = append(bytes, strconv.Itoa(len(value.bulk))...)
	bytes = append(bytes, '\r', '\n')
	bytes = append(bytes, value.bulk...)
	bytes = append(bytes, '\r', '\n')

	return bytes
}

func (value Value) marshalString() []byte {
	var bytes []byte
	bytes = append(bytes, STRING)
	bytes = append(bytes, value.str...)
	bytes = append(bytes, '\r', '\n')

	return bytes
}

func (value Value) marshalInteger() []byte {
	var bytes []byte
	bytes = append(bytes, INTEGER)
	bytes = append(bytes, strconv.Itoa(value.num)...)
	bytes = append(bytes, '\r', '\n')

	return bytes
}

func (value Value) marshalError() []byte {
	var bytes []byte
	bytes = append(bytes, ERROR)
	bytes = append(bytes, value.str...)
	bytes = append(bytes, '\r', '\n')

	return bytes
}

func (value Value) marshalNull() []byte {
	return []byte("$-1\r\n")
}
