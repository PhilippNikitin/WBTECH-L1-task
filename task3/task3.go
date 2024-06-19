package main

import "fmt"

// объявляем функцию для расчета квадрата каждого числа из слайса, подсчета их суммы и записи итогового значения в канал
func sumOfSquares(sum int, nums []int, ch chan int) {
	for _, num := range nums {
		sum += num * num
	}
	// записываем сумму квадратов в sum
	ch <- sum
}

func main() {
	var sumSquares int                    // создаем переменную для подсчета суммы квадратов
	ch := make(chan int)                  // создаем небуферизованный канал
	nums := []int{2, 4, 6, 8, 10}         // инициализируем слайс с указанными в условии значениями
	go sumOfSquares(sumSquares, nums, ch) // запускаем функцию sumOfSquares в асинхронном режиме
	fmt.Println(<-ch)                     // выводим в stdout результат вычислений
}
