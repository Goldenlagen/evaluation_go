package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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
	var requestBody GetSitesRequest
	// buffer := make([]byte, 1024)
	// json.Unmarshal(buffer[GetBody(request)], &requestBody)
	r := io.Read
	err := json.NewDecoder(GetBody(request)).Decode(&requestBody)
	if err != nil {
		fmt.Print("Error decoding request body", err)
	}

	fmt.Print("Body : ", requestBody.query)
}

func PostSites(writer http.ResponseWriter, request *http.Request) {}

func GetFiles(writer http.ResponseWriter, request *http.Request) {}

func GetBody(r *http.Request) []byte {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Panic(err)
	}
	log.Println(string(body))
	return body
}
