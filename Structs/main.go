package main

import "fmt"

func main() {

	fmt.Println("Structs")
	// No inheritance in golang; No super or parent

	clint := User{"Clint", "c@i.com", true, 24}
	fmt.Println(clint)
	fmt.Printf("Clint's details are : %+v \n", clint)
	fmt.Printf("name : %v Email : %v \n", clint.Name, clint.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
