package main

import (
	"fmt"
	"strings"
)

// Lesson 4 Homework

// func main() {
// 	fmt.Println("Hello world!")
// }

// Lesson 5 Homework

var (
	dividend   float64
	devisor    float64
	quotient   float64
	userAnswer string
)

func division(dividend, devisor float64) (float64, error) {
	if devisor == 0 {
		return 0, fmt.Errorf("ZeroDivisionError")
	}
	quotient = dividend / devisor
	return quotient, nil
}

func main() {
	for {
		fmt.Println("Please enter the first number: ")
		_, err := fmt.Scan(&dividend)
		if err != nil {
			fmt.Println("Error: You should enter number")
			continue
		}
		fmt.Println("Please enter the second number: ")
		_, err = fmt.Scan(&devisor)
		if err != nil {
			fmt.Println("Error: You should enter number")
			continue
		}
		divisionResult, zeroDivisionError := division(dividend, devisor)
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
		fmt.Println("Do you want to use another numbers?(y/n)")
		fmt.Scan(&userAnswer)
		if strings.ToLower(userAnswer) == "y" {
			continue
		} else if strings.ToLower(userAnswer) == "n" {
			break
		}
	}
}
