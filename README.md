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
- [ ] DEL
- [x] HSET
- [x] HGET
- [ ] HDEL
- [ ] HGETALL
- [ ] HLEN

## To-Do
- [ ] GoRoutines / Threads - Mutex
- [ ] AOF / Append Only File - Persist Data
