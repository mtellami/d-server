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

## Commands:
- [x] PING
- [x] SET
- [x] GET
- [x] DEL
- [x] HSET
- [x] HGET
- [x] HGETALL
- [x] HDEL

## To-Do
- [x] Multiple Clinets / GoRoutines / Threads - Mutex
- [x] AOF / Append Only File - Data Persistence
