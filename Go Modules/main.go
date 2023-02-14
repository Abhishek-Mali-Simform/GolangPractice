package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello mod in golang")
	greeter()
	router := mux.NewRouter()
	router.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func greeter() {
	fmt.Println("Hey there mod users..")
}

func serveHome(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("<h1>Welcome to Golang Modules</h1>"))
}
