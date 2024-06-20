package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	var t time.Duration // объявляем переменную типа time.Duration, в которой будет храниться установленное пользователем количество секунд - время работы программы
	fmt.Println("Введите время, в течение которого будет работать программа (количество секунд):")
	_, err := fmt.Scanln(&t) // считываем время, в течение которого будет работать программа (в секундах)
	if err != nil {
		fmt.Println("Ошибка при считывании времени:", err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)

	// создаем горутину для непрерывной отправки данных в канал
	go func() {
		var i int
		for {
			ch <- i
			i++
		}
	}()

	// создаем горутину, которая будет одновременно ожидать данные из канала, а также истечения таймаута
	go func() {
		defer wg.Done()
		timeout := time.After(t * time.Second)
		for {
			select {
			case value := <-ch:
				fmt.Println("Получено значение из канала:", value)
			case <-timeout:
				fmt.Println("Время истекло, завершаем работу.")
				return
			}
		}
	}()

	wg.Wait()
}
