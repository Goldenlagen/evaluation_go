package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"evaluation/pkg/protocol"
)

func main() {
	var database protocol.Database
	var defaultSite protocol.Site
	var defaultFile protocol.File

	defaultSite.Id = 1
	defaultSite.HostIp = "10.0.0.1"
	defaultSite.Lastseen = time.Now()
	defaultSite.Domain = "https://test.com"

	defaultFile.Id = 1
	defaultFile.Name = "Accueil"
	defaultFile.SiteId = 1
	defaultFile.Url = "https://test.com/accueil"
	defaultFile.Lastseen = time.Now()

	database.Sites = append(database.Sites, defaultSite)
	database.Files = append(database.Files, defaultFile)

	ln, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	buffer := make([]byte, 1024)
	readchar, _ := conn.Read(buffer)

	fmt.Println(string(buffer[0:readchar]))

	conn.Write([]byte("SERVER RESPONSE"))
	defer conn.Close()
}
