package main

import("fmt"
		"bufio"
		"os"
		"strconv"
	"strings")

func main(){
	fmt.Println("Conversion")
	fmt.Println("Enter rating")
	reader := bufio.NewReader(os.Stdin)
	input,_ := reader.ReadString('\n')
	fmt.Printf("%s", input)
	numRating,err := strconv.ParseFloat(strings.TrimSpace(input),64)
if err != nil{
	fmt.Println(err)
	panic(err)
}
fmt.Println("Rating updated : ",numRating+1)

}