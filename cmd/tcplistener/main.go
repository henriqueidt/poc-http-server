package main

import (
	"fmt"
	"io"
	"net"
	"strings"
)

func readLines(f io.ReadCloser) <-chan string {
	lines := make(chan string)

	go func() {
		defer f.Close()
		defer close(lines)
		
		line := ""
		for {
			b1 := make([]byte, 8)
			n1, err := f.Read(b1)
			if(err != nil) {
				if line != "" {
					lines <- line
				}
				if(err == io.EOF) {
					break
				}
				panic(err)
			}
			
			currContent := string(b1[:n1])
			parts := strings.Split(currContent, "\n")
			
			for i := 0; i < len(parts)-1; i++ {
			lines <- fmt.Sprintf("%s%s", line, parts[i])
			line = ""
		}
		
		line += parts[len(parts)-1]
		
		}
	}()

	return lines;
}

func main() {
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		for line := range readLines(conn) {
			fmt.Println("read:", line)
		}
	}
}