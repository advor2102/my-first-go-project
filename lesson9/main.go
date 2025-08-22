package main

import (
	"fmt"
)

func addValToChan(ch chan int) {
	defer close(ch)
	for i := 0; i <= 20; i++ {
		ch <- i
		fmt.Println("Added: ", i)
	}
}

func readValFromChan(ch chan int) {
	for num := range ch {
		fmt.Println("Printed: ", num)

	}
}
func main() {
	ch := make(chan int, 5)
	go addValToChan(ch)
	readValFromChan(ch)
}
