package main

import "fmt"

// Lesson 4 Homework

// func main() {
// 	fmt.Println("Hello world!")
// }

// Lesson 5 Homework

// var (
// 	dividend   float64
// 	devisor    float64
// 	quotient   float64
// 	userAnswer string
// )

// func division(dividend, devisor float64) (float64, error) {
// 	if devisor == 0 {
// 		return 0, fmt.Errorf("error: you can't divide by zero")
// 	}
// 	quotient = dividend / devisor
// 	return quotient, nil
// }

// func main() {
// 	for {
// 		fmt.Println("Please enter the first number: ")
// 		_, err := fmt.Scan(&dividend)
// 		if err != nil {
// 			fmt.Println("error: you should enter number")
// 			fmt.Println("Do you want to use another numbers?(y/n)")
// 			fmt.Scan(&userAnswer)
// 			if strings.ToLower(userAnswer) == "y" {
// 				continue
// 			} else if strings.ToLower(userAnswer) == "n" {
// 				break
// 			}
// 		}
// 		fmt.Println("Please enter the second number: ")
// 		_, err = fmt.Scan(&devisor)
// 		if err != nil {
// 			fmt.Println("error: you should enter number")
// 			fmt.Println("Do you want to use another numbers?(y/n)")
// 			fmt.Scan(&userAnswer)
// 			if strings.ToLower(userAnswer) == "y" {
// 				continue
// 			} else if strings.ToLower(userAnswer) == "n" {
// 				break
// 			}
// 		}
// 		divisionResult, zeroDivisionError := division(dividend, devisor)
// 		if zeroDivisionError != nil {
// 			fmt.Println(zeroDivisionError)
// 		} else {
// 			fmt.Println(divisionResult)
// 			switch {
// 			case divisionResult > 10:
// 				fmt.Println("Результат большой")
// 			case divisionResult == 1 && divisionResult <= 10:
// 				fmt.Println("Результат средний")
// 			default:
// 				fmt.Println("Результат маленький или ноль")
// 			}
// 		}
// 		fmt.Println("Do you want to use another numbers?(y/n)")
// 		fmt.Scan(&userAnswer)
// 		if strings.ToLower(userAnswer) == "y" {
// 			continue
// 		} else if strings.ToLower(userAnswer) == "n" {
// 			break
// 		}
// 	}
// }

// Lesson 6 Homework

type Person struct {
	Name   string
	Salary int
}

var arr = [5]int{1, 2, 3, 4, 5}
var slice = []int{1, 2, 3, 4, 5}
var ages = map[string]int{
	"Ann":   30,
	"John":  45,
	"Marie": 28,
}
var employee = Person{
	Name:   "Jorge",
	Salary: 1000,
}
var namesList = []string{"Ann", "Bob", "John", "Michael", "Bob", "Ann", "Lisa", "Bob"}

func sum(fTerm int, sTerm int) int {
	sum := fTerm + sTerm
	return sum
}

func doubling(nums []int) []int {
	doubleSlice := make([]int, len(nums))
	for i, v := range nums {
		doubleSlice[i] = v * 2
	}
	return doubleSlice
}

func division(dividend int, devisor int) (int, float64) {
	quotient := dividend / devisor
	remainder := dividend % devisor
	return quotient, float64(remainder)
}

func unite(fSlice []int, sSlice []int) []int {
	fSlice = append(fSlice, sSlice...)
	return fSlice
}

func usingFunc(n int, f func(int) int) int {
	return f(n)
}

func perimeter(a int) int {
	return 2 * a
}

func nameCounter(people []string) map[string]int {
	counter := make(map[string]int)
	for _, v := range people {
		counter[v] += 1
	}
	return counter
}

func main() {
	fmt.Println(sum(5, 6))
	fmt.Println(division(10, 3))
	fmt.Println(arr)
	fmt.Println(slice)
	newSlice := doubling(slice)
	fmt.Println(slice)
	fmt.Println(newSlice)
	fmt.Println(unite(slice, newSlice))
	fmt.Println(ages)
	delete(ages, "Ann")
	value, exists := ages["Ann"]
	// fmt.Println(value, exist)
	if exists {
		fmt.Println(value)
	} else {
		fmt.Println("No data")
	}
	fmt.Println(ages)
	fmt.Println(employee)
	b := usingFunc(5, perimeter)
	fmt.Println(b)
	fmt.Println(nameCounter(namesList))
}
