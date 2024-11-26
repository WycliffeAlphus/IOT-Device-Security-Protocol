package server

import (
	"log"
	"net"
)

// StartServer: Our digital receptionist
func StartServer() {
	// Choose a "phone number" (address)
	address := "localhost:8080"

	// Set up a "phone line" to listen for calls
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Couldn't set up the phone line: %v", err)
	}
	defer listener.Close()

	log.Printf("Server is ready, waiting for connections on %s", address)

	// Forever wait for incoming "calls"
	for {
		// Accept a new connection
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Missed a call: %v", err)
			continue
		}

		// Handle each "caller" in a separate conversation
		go handleClient(conn)
	}
}


