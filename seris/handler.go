package seris

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

var memory = Data{
	mu: sync.RWMutex{},
	hmu: sync.RWMutex{},
	SETs: map[string]string{},
	HSETs: map[string]map[string]string{},
}

var defaultHandlers = map[string]func([]Value) Value {
	"PING": 		ping,
	"SET": 			set,
	"GET": 			get,
	"DEL": 			del,
	"HSET": 		hset,
	"HGET":			hget,
	"HGETALL":	hgetall,
	"HDEL": 		hdel,
}

// HANDLERS
func ping(args []Value) Value {
	return Value{typ: "string", str: "PONG"}
}

// SET VALUE
func set(args []Value) Value {
	if len(args) != 2 {
		 return Value{typ: "error", str: "ERR wrong number of arguments for 'set' command"}
	}

	key := args[0].bulk
	value := args[1].bulk

	memory.mu.Lock()
	memory.SETs[key] = value
	memory.mu.Unlock()

	return Value{typ: "string", str: "OK"}
}

// GET VALUE
func get(args []Value) Value {
	if len(args) != 1 {
		return Value{typ: "error", str: "ERR wrong number of arguments for 'get' command"}
	}

	key := args[0].bulk

	memory.mu.RLock()
	value, ok := memory.SETs[key]
	memory.mu.RUnlock()

	if !ok {
		return Value{typ: "null"}
	}

	return Value{typ: "bulk", bulk: value}
}

// SET HASH VALUE
func hset(args []Value) Value {
	if len(args) != 3 {
		return Value{typ: "error", str: "ERR wrong number of arguments for 'hset' command"}
	}

	hash := args[0].bulk
	key := args[1].bulk
	value := args[2].bulk

	memory.hmu.Lock()
	if _, ok := memory.HSETs[hash]; !ok {
		memory.HSETs[hash] = map[string]string{}
	}

	memory.HSETs[hash][key] = value
	memory.hmu.Unlock()

	return Value{typ: "string", str: "OK"}
}

// GET HASH VALUE
func hget(args []Value) Value {
	if len(args) != 2 {
		return Value{typ: "error", str: "ERR wrong number of arguments for 'hget' command"}
	}

	hash := args[0].bulk
	key := args[1].bulk

	memory.hmu.RLock()
	value, ok := memory.HSETs[hash][key]
	memory.hmu.RUnlock()

	if !ok {
		return Value{typ: "null"}
	}

	return Value{typ: "bulk", bulk: value}
}

// DELETE SINGLE OR MULTIPLE VALUES
func del(args []Value) Value {
	if len(args) == 0 {
		return Value{typ: "error", str: "ERR wrong number of arguments for 'del' command"}
	}

	n := 0

	memory.mu.Lock()
	for i := 0; i < len(args); i++ {
		key := args[i].bulk

		_, ok := memory.SETs[key]
		if ok {
			n += 1
			delete(memory.SETs, key)
		}
	}
	memory.mu.Unlock()

	return Value{typ: "integer", num: n}
}

// GET ALL HASH MAP VALUES
func hgetall(args []Value) Value {
	if len(args) != 1 {
		return Value{typ: "error", str: "ERR wring number of arguments for 'hgetall' command"}
	}

	key := args[0].bulk

	memory.hmu.RLock()
	records, ok := memory.HSETs[key]
	memory.hmu.RUnlock()

	if !ok {
		return Value{typ: "null"}
	}

	var values []Value
	for _key, _value := range records {
		values = append(values, Value{typ: "string", str: _key})
		values = append(values, Value{typ: "string", str: _value})
	}

	return Value{typ: "array", array: values}
}

// DELETE HASH MAP FIELDS
func hdel(args []Value) Value {
	if len(args) < 2 {
		return Value{typ: "error", str: "ERR wrong number of arguments for 'hdel' command"}
	}

	n := 0
	hash := args[0].bulk

	memory.hmu.Lock()
	if _, ok := memory.HSETs[hash]; !ok {
		return Value{typ: "integer", num: n}
	}

	for i := 1; i < len(args); i++ {
		key := args[i].bulk
		if _, ok := memory.HSETs[hash][key]; ok {
			n += 1
			delete(memory.HSETs[hash], key)
		}
	}
	memory.hmu.Unlock()

	return Value{typ: "integer", num: n}
}

// COMMANDS EXECUTION
func execute(value Value) (Value, error) {
	command := strings.ToUpper(value.array[0].bulk)
	args := value.array[1:]

	handler, ok := defaultHandlers[command]
	if !ok {
		return Value{typ: "string", str: ""}, errors.New(fmt.Sprintf("Invalid command: %s", command))
	}

	return handler(args), nil
}
