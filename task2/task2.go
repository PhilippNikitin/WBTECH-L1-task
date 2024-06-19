package main

import (
	"fmt"
	"time"
)

func Square(nums []int, ch chan int) {
	for _, num := range nums {
		ch <- num * num
	}
	close(ch)
}

func main() {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	ch := make(chan int, 5)
	nums := []int{2, 4, 6, 8, 10}
	go Square(nums, ch)
	for square := range ch {
		fmt.Println(square)
	}
}
