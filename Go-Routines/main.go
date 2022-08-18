package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup //pointer
var ends []string
var mut sync.Mutex //pointer

func main() {
	fmt.Println("Go Routines")
	// go greet("Hello")
	// greet("World")

	websiteList := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}
	for _, val := range websiteList {
		// Add to wait group
		go getStatusCode(val)
		wg.Add(1)
	}
	// Wait at the end of main function
	wg.Wait()
	fmt.Println(ends)
}

func greet(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(s)
	}
}

func getStatusCode(endpoint string) {
	// Execute done at end of execution
	defer wg.Done()
	mut.Lock()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOps in endpount")
	}
	ends = append(ends, endpoint)
	fmt.Printf("%d status code for %s \n", res.StatusCode, endpoint)
	mut.Unlock()
}
