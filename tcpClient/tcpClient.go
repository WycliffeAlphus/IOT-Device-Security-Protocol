package tcpClient

import (
	"fmt"
	"net"
)

func Client() {
	// connect to the Server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	// send some data to the server

	_, err = conn.Write([]byte("Hello, server!"))
	if err != nil {
		fmt.Println(err)
		return
	}
	// Close the Connection
	conn.Close()
}
