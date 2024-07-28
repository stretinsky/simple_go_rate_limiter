package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	tokensLimit := 5
	tokens := make(chan int, tokensLimit)

	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		go handleConnection(conn, tokens)
	}
}

func handleConnection(conn net.Conn, tokens chan int) {
	defer conn.Close()

	clientName := conn.RemoteAddr().Network() + " | " + conn.RemoteAddr().String()
	log.Printf("The connection with client %s is open\n", clientName)

	select {
	case tokens <- 1:
		reader := bufio.NewReader(conn)
		
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		log.Printf("Received from the client: %s", line)

		conn.Write([]byte("Hello from server\n"))
		conn.Close()

		log.Printf("The connection with client %s is closed", clientName)

		<- tokens
	default:
		log.Printf("Client %s exceeded rate limit, closing connection", clientName)
		conn.Write([]byte("Rate limit exceeded, connection closed\n"))
		return
	}	
}
