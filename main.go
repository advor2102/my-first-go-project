package main

import "fmt"

// Lesson 4 Homework

// func main() {
// 	fmt.Println("Hello world!")
// }

// Lesson 5 Homework

var (
	dividend int
	devisor  int
	quotient int
)

func divide(dividend, devisor int) (int, error) {
	if devisor == 0 {
		return 0, fmt.Errorf("ZeroDivisionError")
	}
	quotient = dividend / devisor
	return quotient, nil
}

func main() {
	fmt.Println("Please enter the first number: ")
	fmt.Scan(&dividend)
	fmt.Println("Please enter the second number: ")
	fmt.Scan(&devisor)
	divisionResult, zeroDivisionError := divide(dividend, devisor)
	if zeroDivisionError != nil {
		fmt.Println(zeroDivisionError)
	} else {
		fmt.Println(divisionResult)
		switch {
		case divisionResult > 10:
			fmt.Println("Результат большой")
		case divisionResult == 1 && divisionResult <= 10:
			fmt.Println("Результат средний")
		default:
			fmt.Println("Результат маленький или ноль")
		}

	}

}
