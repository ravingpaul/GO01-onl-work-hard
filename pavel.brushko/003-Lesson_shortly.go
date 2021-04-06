package main

import (
	"fmt"
)

var digit1, counter int

func main() {
	fmt.Println("Please input the digit")
	fmt.Scan(&digit1)
	fmt.Println("Yor digit is ", digit1)
	first1()
}

func first1() {
	digit1++
	counter++
	fmt.Println("Yor digit on", counter, "iteration is", digit1)
	if counter < 4 {
		first1()
	}
}
