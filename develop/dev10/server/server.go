package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":5023")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		go telnetConn(conn)
	}
}

func telnetConn(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		// Read text from client
		text := scanner.Text()
		fmt.Printf("Received: %s\n", text)

		// Write response back to client
		response := "You said: " + text + "\n"
		conn.Write([]byte(response))
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error reading from client: %v\n", err)
	}
}
