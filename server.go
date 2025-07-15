package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	// net.Listen interface is used to accept connections
  ln, err := net.Listen("tcp", ":8000")
	if err != nil {
    log.Fatal(err)
  }
  fmt.Println("Listening on port 8000")

	// Call accept to accept incoming connections
	// It returns a net.Conn interface which can be used to read and write data to the connection
	conn, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}

	for {
		// Read data from the connection
		message, err :=  bufio.NewReader(conn).ReadString('\n')
		if err != nil {
				log.Fatal(err)
		}
		fmt.Print("Message Received:", string(message))
		newmessage := strings.ToUpper(message)
		// Write data back to the connection
		conn.Write([]byte(newmessage + "\n"))
	}
}