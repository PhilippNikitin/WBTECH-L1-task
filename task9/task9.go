package main

import "fmt"

func main() {
	numsChan := make(chan int)
	doublesChan := make(chan int)

	// инициализируем слайс с числами
	numbers := []int{1, 3, 5, 8, 9}

	// стартуем горутину, чтобы прочитать все данные из слайса
	go func() {
		for _, num := range numbers {
			numsChan <- num
		}
		close(numsChan)
	}()

	// стартуем горутину, чтобы удвоить каждое значение из канала numsChan
	go func() {
		for num := range numsChan {
			doublesChan <- 2 * num
		}

		close(doublesChan)
	}()

	for double := range doublesChan {
		fmt.Println(double)
	}

}
