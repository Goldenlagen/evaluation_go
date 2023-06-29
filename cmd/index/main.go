package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"evaluation/pkg/protocol"
)

func main() {
	initDatabase()

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

func initDatabase() protocol.Database {
	var database protocol.Database
	var defaultDir protocol.Dir
	var defaultFile protocol.File

	defaultDir.Id = 1
	defaultDir.HostIp = "10.0.0.1"
	defaultDir.Lastseen = time.Now()
	defaultDir.Domain = "https://test.com"

	defaultFile.Id = 1
	defaultFile.Name = "Accueil"
	defaultFile.SiteId = 1
	defaultFile.Url = "https://test.com/accueil"
	defaultFile.Lastseen = time.Now()

	database.Dirs = append(database.Dirs, defaultDir)
	database.Files = append(database.Files, defaultFile)

	return database
}

func handleConnection(conn net.Conn) {
	buffer := make([]byte, 1024)
	readchar, _ := conn.Read(buffer)

	fmt.Println(string(buffer[0:readchar]))

	conn.Write([]byte("SERVER RESPONSE"))
	defer conn.Close()
}

func createSiteRequest(hostIp string, domain string) {
	now := time.Now()
	var newSite protocol.Dir
	newSite.HostIp = hostIp
	newSite.Domain = domain
	newSite.Lastseen = now
}

func createSiteReponse() {

}
