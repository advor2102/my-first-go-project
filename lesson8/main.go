package main

import (
	"fmt"
	"sync"
	"time"
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

//2. Параллельная загрузка

func dFirstFile(wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println("file1 loaded")
	wg.Done()
}

func dSecondFile(wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	fmt.Println("file2 loaded")
	wg.Done()
}

func dFirdFile(wg *sync.WaitGroup) {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("file3 loaded")
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

	wg.Add(1)
	dFirstFile(&wg)

	wg.Add(1)
	dSecondFile(&wg)
	wg.Wait()

	wg.Add(1)
	dFirdFile(&wg)
	wg.Wait()
}
