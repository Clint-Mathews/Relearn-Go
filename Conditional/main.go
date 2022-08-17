package main

import "fmt"

func main() {
	fmt.Println("Conditional")
	count := 10
	var res string
	if count < 10 {
		res = "Greater"
	} else {
		res = "Lesser"
	}
	fmt.Println(res)

	if num := 30; num < 10 {
		fmt.Println("Greater than 10")
	} else {
		fmt.Println("Less than 10")

	}
}
