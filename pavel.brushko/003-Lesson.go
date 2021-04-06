package main

import (
	"fmt"
)

var digit int

func main() {

	fmt.Scan("%d", &digit)
	fmt.Println("Please input the digit", digit)
	fmt.Println("Yor digit is ", digit)
	first()
}

func first() {
	digit++
	fmt.Println("Yor digit on first iteration is ", digit)
	second()
}

func second() {
	digit++
	fmt.Println("Yor digit on second iteration is ", digit)
	third()
}

func third() {
	digit++
	fmt.Println("Yor digit on third iteration is ", digit)
	fourth()
}

func fourth() {
	digit++
	fmt.Println("Yor digit on fourth iteration is ", digit)

}
