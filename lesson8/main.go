package main

import (
	"fmt"
	"sync"
)

// 1. Печать чисел в двух горутинах

func oddN(nums []int, wg *sync.WaitGroup) {
	for _, v := range nums {
		if v%2 == 0 {
			fmt.Println(v)
		}
	}
	wg.Done()
}

func evenN(nums []int, wg *sync.WaitGroup) {
	for _, v := range nums {
		if v%2 == 1 {
			fmt.Println(v)
		}
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	wg.Add(1)
	go oddN(arr, &wg)

	wg.Add(1)
	go evenN(arr, &wg)

	wg.Wait()
}
