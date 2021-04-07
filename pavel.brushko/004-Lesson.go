package main

import "fmt"

func main() {
	var tasknumber int
	fmt.Println("Input number of task")
	fmt.Scan(&tasknumber)

	if tasknumber == 1 {
		task1()
	}
	if tasknumber == 2 {
		task2()
	}

	if tasknumber == 3 {
		task3()
	}

}

func task1() {
	var number int
	fmt.Println("Input digit between 1 and 59")
	fmt.Scan(&number)
	number = number / 15
	if number < 1 {
		fmt.Println("First")
	} else if number < 2 {
		fmt.Println("Second")
	} else if number < 3 {
		fmt.Println("Third")
	} else if number < 4 {
		fmt.Println("Fourth")
	} else {
		fmt.Println("Check the nuber you were inputed")
	}
}

func task2() {
	var a, b int
	fmt.Println("Input a")
	fmt.Scan(&a)
	fmt.Println("Input b")
	fmt.Scan(&b)

	if (a <= 1) && (b >= 3) {
		a = a + b
		fmt.Println(a)
	} else {
		a = a - b
		fmt.Println(a)
	}

}

func task3() {
	var num int
	fmt.Println("Please input a digit between 1 and 4")
	fmt.Scan(&num)
	fmt.Println(season(num))

}

func season(num int) string {
	if num == 1 {
		return "Winter"
	} else if num == 2 {
		return "Spring"
	} else if num == 3 {
		return "Summer"
	} else if num == 4 {
		return "Fall"
	} else {
		return "Check the number"
	}

}
