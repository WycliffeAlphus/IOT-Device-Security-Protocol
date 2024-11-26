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

func handleClient(conn net.Conn) {
	// Always clean up after the conversation
	defer conn.Close()

	log.Printf("New client connected: %s", conn.RemoteAddr())

	// Prepare a message buffer
	buffer := make([]byte, 1024)

	// Keep the conversation going
	for {
		// Read message from client
		n, err := conn.Read(buffer)
		if err != nil {
			log.Printf("Connection error: %v", err)
			return
		}

		// Echo the message back
		message := string(buffer[:n])
		log.Printf("Received: %s", message)
		conn.Write([]byte("Server says: " + message))
	}
}
