package main

import (
	"time"

	"tcpIp/client"
	"tcpIp/server"
)

func main() {
	// Start server in background
	go func() {
		server.StartServer()
	}()

	// Give server a moment to start
	time.Sleep(time.Second)

	// Launch client
	client.StartClient()
}
