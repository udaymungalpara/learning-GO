package main

import (
	"fmt"
	"sort"
)

func main()  {
	
	fmt.Println("Slices on the GOlang: ")

	var bois = []string{"Uday","Raj","Rahul"}

	fmt.Printf("The type is :%T\n",bois)
	
	bois = append(bois, "sujal","Nikunj")
	
	fmt.Println(bois)

	bois = append(bois[2:5])
	fmt.Println(bois)

	Numbers := make([]int, 5)

	
	Numbers[0] = 767
	Numbers[1] = 567
	Numbers[2] = 987
	Numbers[3] = 231
	Numbers[4] = 546

	fmt.Println(Numbers)

	sort.Ints(Numbers)
	fmt.Println(Numbers)
}