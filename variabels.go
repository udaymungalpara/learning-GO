package main

import "fmt"

const loginuser string = "udaymungalpara"

func main() {

	var name string = "Uday"
	fmt.Println(name)
	fmt.Printf("The type of variable is:%T \n", name)

	var isgood bool = false
	fmt.Println(isgood)
	fmt.Printf("The type of variable is:%T \n", isgood)

	var num int = 500
	fmt.Println(num)
	fmt.Printf("The type of variable is:%T \n", num)

	var floa float64 = 5000.66462
	fmt.Println(floa)
	fmt.Printf("The type of variable is:%T \n", floa)

	var floa2 float32 = 500.473758
	fmt.Printf("The typee of variable is :%T\n", floa2)
	fmt.Println(floa2)

	fmt.Println(loginuser)
	fmt.Printf("The type of variable is  :%T\n", loginuser)
}
