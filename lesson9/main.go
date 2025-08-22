package main

import (
	"fmt"
)

func addValToChan(ch chan int) {
	for i := 0; i <= 20; i++ {
		ch <- i
	}
}

func readValFromChan(ch chan int) {
	for num := range ch {
		fmt.Println(num)
	}
}
func main() {
	ch := make(chan int, 5)
	go addValToChan(ch)
	readValFromChan(ch)
}
