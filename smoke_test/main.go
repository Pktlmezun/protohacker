package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	message, err := reader.ReadString('\n')
	if err != nil {
		log.Println("Error reading string from conn: ", err)
	}
	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Println("Error writing to conn: ", err)
	}
	fmt.Println(conn)
}
func main() {
	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatal("doesn't work tcp connection")
	}
	defer listener.Close()

	for {

		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting conn: ", err)
			continue
		}
		go handleConnection(conn)
	}
}
