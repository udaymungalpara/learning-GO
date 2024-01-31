package main

import "fmt"

func main() {

	fmt.Println("Array In golag:")

	var bois [5]string

	
	bois[0]  = "uday"
	bois[2]  = "sagar"
	bois[4]  = "bhautik"

	fmt.Println("The boys are:",bois)
	fmt.Println("The length of array is:",len(bois))

	var girls [3]string

	girls[1]="bhummi"
	girls[2]="nikuu"

	fmt.Println("The boys are:",girls)
	fmt.Println("The length of array is:",len(girls))

}
