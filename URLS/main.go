package main

import (
	"fmt"
	"net/url"
)

const myURL string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=sdsds"

func main() {
	fmt.Println("URLs")
	fmt.Println(myURL)

	// parsing
	result, _ := url.Parse(myURL)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	pqparams := result.Query()

	fmt.Printf("Type %T \n", pqparams)
	for _, values := range pqparams {
		fmt.Println("Param is ", values)
	}

	partsOfURL := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "tutcss",
		RawQuery: "user=hitesh",
	}

	anotherURL := partsOfURL.String()
	fmt.Println(anotherURL)

}
