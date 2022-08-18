package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Get Request")
	performGetRequest()
}

func performGetRequest() {
	const url = "http://localhost:9088/get"
	byteData, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer byteData.Body.Close()
	fmt.Println("Status Code", byteData.Status)
	fmt.Println("Content length", byteData.ContentLength)

	data, err := ioutil.ReadAll(byteData.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Byte Data", data)
	fmt.Println("Data", string(data))

	var responseString strings.Builder

	byteCount, _ := responseString.Write(data)

	fmt.Println("Byte Count", byteCount)
	fmt.Println("Data", responseString.String())

}
