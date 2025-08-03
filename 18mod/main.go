package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("mod in Gooooooooooooooo!")
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods(("GET"))

	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {

}

func serveHome(w http.ResponseWriter, rr *http.Request) {
	w.Write([]byte("hello from go backend Soham"))
}
