package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Files")

	content := "Hello World"
	fileName := "./file.txt"
	file, err := os.Create(fileName)

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	checkNilError(err)

	fmt.Println("Length is ", length)
	defer file.Close()
	readFile(fileName)
}

func readFile(fileName string) {
	byteData, err := os.ReadFile(fileName)
	checkNilError(err)

	fmt.Println("Byte data ", byteData)
	fmt.Println("Data ", string(byteData))

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
