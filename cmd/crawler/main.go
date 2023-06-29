package main

import (
	"fmt"
	"log"
	"net"
)

func crawler(conn net.Conn) {
	conn.Write([]byte("This is Crawler"))

	buffer := make([]byte, 1024)
	readchar, readerr := conn.Read(buffer)

	if readerr != nil {
		log.Fatal(readerr)
	} else {
		fmt.Println(string(buffer[0:readchar]))
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	crawler(conn)
}
