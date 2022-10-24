package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Hello from UserInput"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("input a string: ")
	// comma ok / error ok
	// input, _ --> try, catch

	input, _ := reader.ReadString('\n')
	fmt.Println("user entered: ", input)

}
