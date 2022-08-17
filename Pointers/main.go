package main

import "fmt"

func main() {
	fmt.Println("Pointers")
	// var ptr *int
	// fmt.Println("Value of pointer is", ptr)

	myNum := 10
	var ptr *int = &myNum
	fmt.Println("Ptr ", ptr)
	fmt.Println("Ptr ", *ptr)
	*ptr = *ptr * 2
	fmt.Println("Ptr ", *ptr)

}
