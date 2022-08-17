package main

import "fmt"

func main() {
	fmt.Println("Array")
	var fruitList [5]string
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list", len(fruitList), fruitList)

	var vegList = [4]string{"Potato", "Beans", "Mushroom"}
	fmt.Println("Veggie list : ", len(vegList), vegList)
}
