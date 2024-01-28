package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("hello")

	reader := bufio.NewReader(os.Stdin)
	 fmt.Println("enter the price:")

	 input, _:=reader.ReadString('\n')

	 price, err:=strconv.ParseFloat(strings.TrimSpace(input),64)

	 if err != nil{

		fmt.Println("error:",err)
	 }else{
		fmt.Println("The tax of product is:",price*0.18)
	 }

}