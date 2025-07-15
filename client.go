package main

import (
    "bufio"
    "fmt"
    "log"
    "net"
    "os"
)

func main() {
	// net.Dial interface is used to connect to the server
	// It returns a net.Conn interface which can be used to read and write data to the connection
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
			log.Fatal(err)
	}
	for {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Text to send: ")
			text, _ := reader.ReadString('\n')
			// Write data to the connection
			fmt.Fprintf(conn, text + "\n")
			// Read data back from the connection
			message, _ := bufio.NewReader(conn).ReadString('\n')
			fmt.Print("Message from server: "+message)
	}
}