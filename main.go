package main

import (
	"log"

	"iotsec/server"
)

func main() {
	log.Println("Starting Iot Security Protocol Server...")
	server.StartServer()
}
