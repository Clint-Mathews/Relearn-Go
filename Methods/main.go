package main

import "fmt"

func main() {

	fmt.Println("Structs")
	// No inheritance in golang; No super or parent

	clint := User{"Clint", "c@i.com", true, 24}
	fmt.Println(clint)
	fmt.Printf("Clint's details are : %+v \n", clint)
	clint.GetStatus()
	clint.UpdateEmail("kk@y.com")
	fmt.Printf("name : %v Email : %v \n", clint.Name, clint.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Status", u.Status)
}

func (u User) UpdateEmail(email string) {
	u.Email = email
	fmt.Println("Email is ", u.Email)
}
