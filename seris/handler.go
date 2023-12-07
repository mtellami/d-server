package seris

var memory = Data{
	SETs: map[string]string{},
	HSETs: map[string]map[string]string{},
}

var defaultHandlers = map[string]func([]Value) Value {
	"PING": ping,
	"SET": set,
	"GET": get,
	"HSET": hset,
	"HGET": hget,
}

func ping(args []Value) Value {
	return Value{typ: "string", str: "PONG"}
}

func set(args []Value) Value {
	if len(args) != 2 {
		 return Value{typ: "error", str: "ERR wrong number of arguments for 'set' command"}
	}

	key := args[0].bulk
	value := args[1].bulk

	memory.SETs[key] = value

	return Value{typ: "string", str: "OK"}
}

func get(args []Value) Value {
	if len(args) != 1 {
		return Value{typ: "error", str: "ERR wrong number of arguments for 'get' command"}
	}

	key := args[0].bulk

	value, ok := memory.SETs[key]
	if !ok {
		return Value{typ: "null"}
	}

	return Value{typ: "bulk", bulk: value}
}

func hset(args []Value) Value {
	if len(args) != 3 {
		return Value{typ: "error", str: "ERR wrong number of arguments for 'hset' command"}
	}

	hash := args[0].bulk
	key := args[1].bulk
	value := args[2].bulk

	if _, ok := memory.HSETs[hash]; !ok {
		memory.HSETs[hash] = map[string]string{}
	}

	memory.HSETs[hash][key] = value

	return Value{typ: "string", str: "OK"}
}

func hget(args []Value) Value {
	if len(args) != 2 {
		return Value{typ: "error", str: "ERR wrong number of arguments for 'hget' command"}
	}

	hash := args[0].bulk
	key := args[1].bulk

	value, ok := memory.HSETs[hash][key]
	if !ok {
		return Value{typ: "null"}
	}

	return Value{typ: "bulk", bulk: value}
}

