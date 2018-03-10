package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
)

func main() {

    var port int
	log.Println("Please Enter a port number: ")
    _, err := fmt.Scanf("%d", &port)

	l, err := net.Listen("tcp", ":" + strconv.Itoa(port))
	if err != nil {
		log.Panicln(err)
	}
	log.Println("Listening to connections on port", strconv.Itoa(port))
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Panicln(err)
		}

		go handleRequest(conn)
	}
}


func handleRequest(conn net.Conn) {
	log.Println("Accepted new connection.")
	defer conn.Close()
	defer log.Println("Closed connection.")

	for {
		buf := make([]byte, 1024)
		size, err := conn.Read(buf)
		if err != nil {
			return
		}
		data := buf[:size]
		log.Println("Data read from connection", data)
	}
	// TODO: add JSON Parsing and marshling of network data for data stored in buff
	
	//log.Println("Data is in a json format")
}
