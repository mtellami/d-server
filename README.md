# SERIS / DATABASE SERVER
Seris is database server written in Golang inspired by Redis using serialization protocol specification (RESP).

```go
                    .__        
  ______ ___________|__| ______
 /  ___// __ \_  __ \  |/  ___/
 \___ \\  ___/|  | \/  |\___ \ 
/____  >\___  >__|  |__/____  >
     \/     \/              \/ 
```

## RUN
```zsh
go run main.go
```

## Build
```zsh
go build -o resis main.go
```

## Usage
- Initialize seris server
```go
package main

import (
	"fmt"
	"seris/seris"
)

func main() {

	server, err := seris.NewServer(&seris.Config{
		Port: 6379,
		EnableAof: true,
		AofFile: "database.aof",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	server.Listen()
}
```

- Connect to the server using a redis client
```zsh
$ redis-cli
```

## Data Persistence
- Enable / Desable data save on dist
```go
Config {
    EnableAof: false
    ...
}
```


## Supported Commands:
- [x] PING
- [x] SET
- [x] GET
- [x] DEL
- [x] HSET
- [x] HGET
- [x] HGETALL
- [ ] HLEN
- [x] HDEL
