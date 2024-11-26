package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func StartClient() {
	// Dial the server (like making a phone call)
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Printf("Could not connect to server: %v", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server. Start chatting!")

	// Read user input and send to server
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		message := scanner.Text()
		if message == "exit" {
			break
		}

		// Send message to server
		conn.Write([]byte(message + "\n"))

		// Wait for server's response
		response := make([]byte, 1024)
		n, err := conn.Read(response)
		if err != nil {
			log.Printf("Error receiving response: %v", err)
			break
		}

		fmt.Println(string(response[:n]))
	}
}
