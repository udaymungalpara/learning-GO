package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	//var taxamount int

	fmt.Println("Enter the rate of product:-")

	reader := bufio.NewReader(os.Stdin)

	rate, _ := reader.ReadString('\n')

	fmt.Println("The amout before tax:", rate)

	amount, err := strconv.Atoi(rate)

	/*if err != nil {
		fmt.Println(err)
	} else {

		tax := amount % 10
		taxamount := amount + tax
	}
	fmt.Println("The rate after Tax: ", taxamount)
	fmt.Printf("%v", taxamount)*/

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("adding 10:", amount+10)
	}

}
