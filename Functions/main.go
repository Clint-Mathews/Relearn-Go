package main

import "fmt"

func main() {
	fmt.Println("Functions")
	greeet()
	result := add(3, 5)
	fmt.Println("Result is ", result)
	values := []int{1, 2, 3, 4, 6, 7, 3, 4}
	fmt.Println("Result of sum slice ", addSlice(values...))
	fmt.Println("Result of sum slice ", addSlice(2, 3, 4, 21, 212))
	i, v := checkSlice(2, 3, 4, 21, 10, 212)
	fmt.Println("Check if 10 exist index ", i, "value", v)
}

func greeet() {
	fmt.Println("Greet")
}

func add(val1 int, val2 int) int {
	return val1 + val2
}

func addSlice(values ...int) int {
	tot := 0
	for _, i := range values {
		tot += i
	}
	return tot
}

func checkSlice(values ...int) (int, int) {
	for i, v := range values {
		if v == 10 {
			return i, v
		}
	}
	return 0, 0
}
