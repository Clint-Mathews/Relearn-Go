package main

import "fmt"

func main() {
	fmt.Println("Loops")
	days := []string{"Monday", "Tuesday", "Wednesday", "Friday"}

	fmt.Println(days)

	for day := 0; day < len(days); day++ {
		fmt.Println(days[day])
	}

	for day := range days {
		fmt.Println(days[day])
	}

	for index, day := range days {
		fmt.Println(index, day)
	}

	customVal := 1

	for customVal < 10 {
		if customVal == 5 {
			customVal++
			goto lco
			continue
		}
		fmt.Println("Value is ", customVal)

		customVal++
	}

lco:
	fmt.Println("Jumped")
}
