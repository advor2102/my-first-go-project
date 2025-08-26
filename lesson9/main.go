package main

import (
	"fmt"
	"sync"
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
		time.Sleep(2000 * time.Millisecond)
	}
}

func sendToSecCahnel(ch2 chan string) {
	defer close(ch2)
	for i := 1; i <= 20; i++ {
		ch2 <- fmt.Sprintf("ch2: %d", i)
		time.Sleep(2000 * time.Millisecond)
	}
}

func loadData(resultCh chan<- string) {
	fmt.Println("Loading data...")
	time.Sleep(5 * time.Second)
	resultCh <- "Data loaded successful"
}

func wGenerator(w int, results chan<- string) {
	msg := fmt.Sprintf("Worker %d", w)
	results <- msg
}

func collect(n int) []string {
	results := make(chan string, n)
	var wg sync.WaitGroup
	var collector []string
	wg.Add(n)
	for i := 1; i <= n; i++ {
		go func(wID int) {
			defer wg.Done()
			wGenerator(wID, results)
		}(i)
	}
	wg.Wait()
	close(results)
	for r := range results {
		collector = append(collector, r)
	}
	return collector
}

func main() {
	// ch := make(chan int, 5)
	// ch1 := make(chan string)
	// ch2 := make(chan string)
	// go addValToChan(ch)
	// readValFromChan(ch)
	// go sendToFirstCahnel(ch1)
	// go sendToSecCahnel(ch2)
	// timer := time.NewTimer(1500 * time.Millisecond)
	// for ch1 != nil || ch2 != nil {
	// 	select {
	// 	case massage, ok := <-ch1:
	// 		if !ok {
	// 			ch1 = nil
	// 			fmt.Println("ch1 closed")
	// 			continue
	// 		}
	// 		fmt.Println("Received:", massage)
	// 		if !timer.Stop() {
	// 			select {
	// 			case <-timer.C:
	// 			default:
	// 			}
	// 		}
	// 		timer.Reset(1500 * time.Millisecond)
	// 	case massage, ok := <-ch2:
	// 		if !ok {
	// 			ch2 = nil
	// 			fmt.Println("ch2 closed")
	// 			continue
	// 		}
	// 		fmt.Println("Received:", massage)
	// 		if !timer.Stop() {
	// 			select {
	// 			case <-timer.C:
	// 			default:
	// 			}
	// 		}
	// 		timer.Reset(1500 * time.Millisecond)
	// 	case <-timer.C:
	// 		fmt.Println("Timeout: no data received for 1.5 seconds")
	// 		return
	// 	}
	// }
	// resultCh := make(chan string)
	// go loadData(resultCh)
	// fmt.Println("Waiting...")
	// result := <-resultCh
	// fmt.Println("Result:", result)
	results := collect(5)
	fmt.Println("Results:")
	for r := range results {
		fmt.Println(r)
	}
}
