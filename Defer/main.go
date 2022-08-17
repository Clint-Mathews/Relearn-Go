package main

import "fmt"

func main() {
	defer fmt.Println("Defer Hello World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello World")
	myDefer()

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
