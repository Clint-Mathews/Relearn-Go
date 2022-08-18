package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Clint-Mathews/ReLearnGo/MongoDbAPI/41/router"
)

func main() {
	r := router.Router()
	fmt.Println("MongoDb - API")
	fmt.Println("Server starting")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening on port 4000 ...")
}
