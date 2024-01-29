package main

import "fmt"

func main()  {
	

	var ptr *int

	num := 25

	ptr = &num

	fmt.Println("the address of num  :", ptr)
	fmt.Println("the value of num :", *ptr)
	fmt.Println("another method of address ", &ptr)

	
}