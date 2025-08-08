package main

import "fmt"

// Lesson 4 Homework

// func main() {
// 	fmt.Println("Hello world!")
// }

// Lesson 5 Homework

var (
	devidend int
	devisor  int
	quotient int
)

func main() {
	fmt.Println("Please enter the first number: ")
	fmt.Scan(&devidend)
	fmt.Println("Please enter the second number: ")
	fmt.Scan(&devisor)
	if devisor == 0 {
		fmt.Println("ZeroDivisionError")
	} else {
		quotient = devidend / devisor
		fmt.Println(quotient)
	}
}
