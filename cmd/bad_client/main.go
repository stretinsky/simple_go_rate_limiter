package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	delim := "\n"
	message := []string{"H", "e", "l", "l", "o", " ", "f", "r", "o", "m", " ", "b", "a", "d", " ", "c", "l", "i", "e", "n", "t", delim}

	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	for _, v := range message {
		conn.Write([]byte(v))
		log.Printf("Write: %s", v)
		time.Sleep(time.Millisecond * 500)
	}

	reader := bufio.NewReader(conn)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Printf("Received from the server: %s", line)


	conn.Close()
}
