package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	weclome := "Welcome ! User input"
	fmt.Println(weclome)
	fmt.Println("Enter line")
	reader := bufio.NewReader(os.Stdin)

	// Comma OK || Error Ok Syntax
	input, _ := reader.ReadString('\n')

	fmt.Printf("input : %s \n", input)
	fmt.Printf("Type of reader is : %T \n", reader)
	fmt.Printf("Type of input is : %T \n", input)

}
