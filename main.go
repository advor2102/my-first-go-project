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

func devide(devidend, devisor int) (int, error) {
	if devisor == 0 {
		return 0, fmt.Errorf("ZeroDivisionError")
	}
	quotient = devidend / devisor
	return quotient, nil
}

func main() {
	fmt.Println("Please enter the first number: ")
	fmt.Scan(&devidend)
	fmt.Println("Please enter the second number: ")
	fmt.Scan(&devisor)
	fmt.Println(devide(devidend, devisor))
}
