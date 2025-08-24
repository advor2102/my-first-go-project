package main

import (
	"fmt"
	"time"
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

func sendToFirstCahnel(ch1 chan string) {
	defer close(ch1)
	for i := 1; i <= 10; i++ {
		ch1 <- fmt.Sprintf("ch1: %d", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func sendToSecCahnel(ch2 chan string) {
	defer close(ch2)
	for i := 1; i <= 20; i++ {
		ch2 <- fmt.Sprintf("ch2: %d", i)
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	ch := make(chan int, 5)
	ch1 := make(chan string)
	ch2 := make(chan string)
	go addValToChan(ch)
	readValFromChan(ch)
	go sendToFirstCahnel(ch1)
	go sendToSecCahnel(ch2)
	for ch1 != nil || ch2 != nil {
		select {
		case massage, ok := <-ch1:
			if !ok {
				ch1 = nil
				fmt.Println("ch1 closed")
				continue
			}
			fmt.Println("Received:", massage)
		case massage, ok := <-ch2:
			if !ok {
				ch2 = nil
				fmt.Println("ch2 closed")
				continue
			}
			fmt.Println("Received:", massage)
		}
	}
}
