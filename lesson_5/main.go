package main

import "fmt"

var (
	first_number  int
	second_number int
)

func main() {
	fmt.Println("Please enter the first number: ")
	fmt.Scan(&first_number)
	fmt.Println("Please enter the second number: ")
	fmt.Scan(&second_number)
	fmt.Println(first_number, second_number)
}
