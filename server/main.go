package main

import (
	"log"

	"github.com/lzhz/disbeauty"
)

type TestStruct struct {
	x int
	Y int
}

func main() {
	server, err := disbeauty.NewServer("127.0.0.1:9090", 1024)
	disbeauty.RegisterData("server", server)
	if err != nil {
		log.Fatalln(err)
	}
	server.Serve()
}
