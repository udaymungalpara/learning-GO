package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	Reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your Age:")

	//comma okay

	input, _ := Reader.ReadString('\n')
	fmt.Println("Your age is:",input)

}