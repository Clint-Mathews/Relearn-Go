package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("Math-Crypto-Random")

	// var num1 int = 10
	// var num2 float64 = 9
	// fmt.Println("Sum : ", num1+int(num2))

	// Random number from math
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println("Random number from math", rand.Intn(5)+1)

	// Random number from crypto
	randomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println("Random number from crypto", randomNum)

}
