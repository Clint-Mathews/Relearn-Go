package main

import "fmt"

func main() {
	fmt.Println("Maps")

	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["RD"] = "Ruby"
	languages["Py"] = "Python"

	fmt.Println("List of all languages : ", languages)
	fmt.Println("JS : ", languages["JS"])

	delete(languages, "RD")
	fmt.Println("List of all languages : ", languages)

	// Loops
	for key, value := range languages {
		fmt.Printf("Key %v , value is :%v \n", key, value)
	}

}
