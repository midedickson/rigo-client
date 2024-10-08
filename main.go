package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	host := os.Args[1]
	port := os.Args[2]
	fmt.Printf("Connecting to Rigo server at %s:%s... 🛫\n", host, port)

	// Connect to the server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Connected to Rigo server at %s:%s ✅\n", host, port)

	defer closeConnection(conn)
	for {
		// Create a new reader from standard input
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("%s:%s> ", host, port)

		// Read until newline is encountered
		commandInput, _ := reader.ReadString('\n')
		data := []byte(fmt.Sprintf("%s\n", commandInput))
		_, err = conn.Write(data)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Wait for a response
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Printf("Error reading response: %v\n", err)
			return
		}
		fmt.Printf("Received response:%s \n", string(buffer[:n]))
	}
}

func closeConnection(conn net.Conn) {
	fmt.Println("Closing the connection to Rigo server ❌")
	conn.Close()
}
