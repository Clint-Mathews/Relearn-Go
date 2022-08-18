package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello MOD in go")
	greet()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4010", r))
}

func greet() {
	fmt.Println("Hello MOD users")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello</h1>"))
}
