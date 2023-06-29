package main

import (
	"encoding/json"
	"evaluation/pkg/protocol"
	"fmt"
	"log"
	"net"
	"time"
)

func crawler(conn net.Conn) {
	var newData protocol.CreateFileRequest
	newData.Command = "createfile"
	newData.File.Id = 1
	newData.File.Lastseen = time.Now()
	newData.File.Name = "FilmAVATAR.mp4"
	newData.File.SiteId = 1
	newData.File.Url = "https://filmdownload/FilmAVATAR.mp4"

	jsonstring, _ := json.Marshal(newData)

	conn.Write(jsonstring)

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

// func handleCrawlerAction(action string, conn net.Conn) {
// 	var
// }
