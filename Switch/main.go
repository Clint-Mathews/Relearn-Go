package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch")
	rand.Seed(time.Now().UnixNano())
	diceNum := rand.Intn(6) + 1
	fmt.Println("Value of dice is ", diceNum)

	switch diceNum {
	case 1:
		fmt.Println("Dice value is 1, Move 1 position")
		fallthrough // Prints both 1 & 2 if we get 1
	case 2:
		fmt.Println("Dice value is 2, Move 2 position")
	case 3:
		fmt.Println("Dice value is 3, Move 3 position")
	case 4:
		fmt.Println("Dice value is 4, Move 4 position")
	case 5:
		fmt.Println("Dice value is 5, Move 5 position")
	case 6:
		fmt.Println("Dice value is 6, Move 6 position")

	default:
		fmt.Println("Wrong roll!")
	}

}
