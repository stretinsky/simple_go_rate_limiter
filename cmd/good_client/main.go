package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	delim := "\n"
	message := "Hello from good client" + delim

	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	conn.Write([]byte(message))
	log.Printf("Write: %s", message)

	reader := bufio.NewReader(conn)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Printf("Received from the server: %s", line)


	conn.Close()
}
