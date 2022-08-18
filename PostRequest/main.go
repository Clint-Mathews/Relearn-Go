package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Post Request")
	// generatePostRequest()
	performPostFormRequest()
}

func generatePostRequest() {
	const url = "http://localhost:9088/post"
	body := strings.NewReader(`
		{
			"Name":"Clint",
			"Age":"23"
		}
	`)
	fmt.Println("Request data", body)

	res, err := http.Post(url, "application/json", body)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Response data", string(data))

}

func performPostFormRequest() {
	const apiUrl = "http://localhost:9088/postForm"

	// Form Data

	data := url.Values{}
	data.Add("Name", "Clint")
	data.Add("Age", "23")

	res, err := http.PostForm(apiUrl, data)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	rsponseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(rsponseBody))

}
