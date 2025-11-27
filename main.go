package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("./messages.txt")
	if(err != nil) {
		panic(err)
	}

	line := ""
	for {
		b1 := make([]byte, 8)
		n1, err := f.Read(b1)
		if(err != nil) {
			if line != "" {
				fmt.Printf("read: %s\n", line)
				line = ""
			}
			if(err == io.EOF) {
				break
			}
			panic(err)
		}

		currContent := string(b1[:n1])
		parts := strings.Split(currContent, "\n")

		for i := 0; i < len(parts)-1; i++ {
			fmt.Printf("read: %s%s\n", line, parts[i])
			line = ""
		}

		line += parts[len(parts)-1]

	}

	// net.Dial interface is used to connect to the server
	// It returns a net.Conn interface which can be used to read and write data to the connection
	// conn, err := net.Dial("tcp", "localhost:8000")
	// if err != nil {
	// 		log.Fatal(err)
	// }
	// for {
	// 		reader := bufio.NewReader(os.Stdin)
	// 		fmt.Print("Text to send: ")
	// 		text, _ := reader.ReadString('\n')
	// 		// Write data to the connection
	// 		fmt.Fprintf(conn, text + "\n")
	// 		// Read data back from the connection
	// 		message, _ := bufio.NewReader(conn).ReadString('\n')
	// 		fmt.Print("Message from server: "+message)
	// }
}