package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Web Requests")

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response of type %T \n", res)
	defer res.Body.Close() // Caller's responsibility to close the connection

	dataBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	content := string(dataBytes)
	fmt.Println("Response", content)

}
