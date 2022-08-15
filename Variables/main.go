package main

import "fmt"

var jwtToken int = 1212       //local
const Token string = "Helllo" //Public

func main() {
	var user string = "Clint"
	fmt.Println("Variables")
	fmt.Println(user)
	fmt.Printf("Varaiable %s is of type %T \n", user, user)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Varaiable %t is of type %T \n", isLoggedIn, isLoggedIn)

	var smallFloat float32 = 255.121
	fmt.Println(smallFloat)
	fmt.Printf("Variabel %f is of type %T \n", smallFloat, smallFloat)

	var anotheInt int
	fmt.Println(anotheInt)
	fmt.Printf("Variabel %d is of type %T \n", anotheInt, anotheInt)

	// Implict type
	var website = "Hello"
	fmt.Println(website)
	// website = 1
	// fmt.Println(website)

	// No var style
	noOfUsers := 10
	fmt.Println(noOfUsers)
	fmt.Printf("Variable %d is of type %T \n", noOfUsers, noOfUsers)

	fmt.Printf("Variable %s is of type %T \n", Token, Token)

}
