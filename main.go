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
