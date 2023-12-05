package main

import (
	"seris/seris"
)

func main() {

	server := seris.NewServer(&seris.Config{
		Port: 6379,
		EnableAof: true,
		AofFile: "database.aof",
	})

	server.Listen()
}
