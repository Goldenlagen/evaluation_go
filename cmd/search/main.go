package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
)

type GetSitesRequest struct {
	query  string
	domain string
}

func search(conn net.Conn) {
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
	runServer()

	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		log.Fatal(err)
	}

	search(conn)

}

func runServer() {
	router := mux.NewRouter()

	router.HandleFunc("/sites", GetSites).Methods("GET")
	router.HandleFunc("/sites", PostSites).Methods("POST")
	router.HandleFunc("/files", GetFiles).Methods("GET")

	err := http.ListenAndServe(":8081", router)
	fmt.Println("Server running ...")
	if err != nil {
		log.Fatalln("Http server error", err)
	}
}

func GetSites(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	requestBody := GetSitesRequest{}
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		fmt.Println("Error decoding request body")
	}

	fmt.Print("Body : ", requestBody)
}

func PostSites(writer http.ResponseWriter, request *http.Request) {}

func GetFiles(writer http.ResponseWriter, request *http.Request) {}
