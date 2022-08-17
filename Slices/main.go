package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices")

	var fruitList = []string{}
	fmt.Printf("Fruit list type %T", fruitList)
	fruitList = append(fruitList, "Apple")
	fruitList = append(fruitList, "Mongo")
	fruitList = append(fruitList, "Banana")
	fmt.Println("Fruit list : ", fruitList)
	fruitList = append(fruitList[1:])
	fmt.Println("Fruit list : ", fruitList)

	fruitList = append(fruitList[1:2]) // End index is non inclusive
	fmt.Println("Fruit list : ", fruitList)

	highscores := make([]int, 4)
	highscores[0] = 10
	highscores[1] = 20
	highscores[2] = 30
	highscores[3] = 40
	// highscores[4] = 50 Index out of bound

	highscores = append(highscores, 12, 12, 12, 12, 1, 21)
	// Appends ands add the new items into slice

	fmt.Println(highscores)
	fmt.Println(sort.IntsAreSorted(highscores))

	sort.Ints(highscores)

	fmt.Println(highscores)
	fmt.Println(sort.IntsAreSorted(highscores))

	// How to remove a value from slices based on index

	var courses = []string{"reactjs", "python", "swift", "go"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
