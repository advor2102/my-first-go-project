package main

import "fmt"

// Lesson 4 Homework

// func main() {
// 	fmt.Println("Hello world!")
// }

// Lesson 5 Homework

var (
	first_number  int
	second_number int
)

func main() {
	fmt.Println("Please enter the first number: ")
	fmt.Scan(&first_number)
	fmt.Println("Please enter the second number: ")
	fmt.Scan(&second_number)
	if second_number == 0 {
		fmt.Println("ZeroDivisionError")
	} else {
		fmt.Println(first_number / second_number)
	}
}
