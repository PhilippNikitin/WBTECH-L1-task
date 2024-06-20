package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
)

// объявляем функцию по созданию воркера worker

// Параметры

// id - уникальный номер воркера, jobs - главный поток, откуда воркер будет получать задания,
// wg - указатель на sync.WaitGroup, для завершения работы воркера в случае нажатия Ctrl + C.
// Логика функции заключается в том, что воркер читает данные из главного потока (канала jobs),
// и выводит их в stdout.
// sync.WaitGroup - это структура, которая используется для синхронизации выполнения нескольких горутин.
// Она отслеживает, сколько горутин еще выполняются, и позволяет основной горутине дождаться завершения всех остальных.
func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Println("worker", id, "прочитал и вывел в stdout данные из главного потока:", j)
	}
}

// объявляем функцию по созданию заданного количества воркеров

// параметры

// number - желаемое количество воркеров, jobs - главный поток, wg - указатель на sync.WaitGroup
func spawnWorkers(number int, jobs chan int, wg *sync.WaitGroup) {
	for w := 1; w <= number; w++ {
		wg.Add(1) // увеличиваем счетчик в WatGroup на 1. Добавляем еще одну горутину, которую основная горутина должна дождаться.
		go worker(w, jobs, wg)
	}
}

// Объявляем функцию Loop для обеспечения постоянной записи данных в главный канал. Для остановки используется вспомогательные канал done
// для канала done мы используем struct{} , чтобы сообщить, что событие произошло, но не передаем никакого значения.
func Loop(n int, ch chan int, done <-chan struct{}) {
	for i := 0; ; i++ {
		select {
		case ch <- i:
		case <-done:
			return
		}
	}
}

func main() {
	var n int // объявляем переменную, в которой будет храниться установленное пользователем количество воркеров с stdin
	fmt.Println("Укажите желаемое число воркеров:")
	_, err := fmt.Scanln(&n) // считываем количестов воркеров, которое хочет запустить пользователь
	if err != nil {
		fmt.Println("Ошибка при считывании числа воркеров:", err)
		return
	}
	jobs := make(chan int, n) // объявляем главный канал jobs, который будет иметь емкость, равную заданному количеству воркеров
	var wg sync.WaitGroup
	done := make(chan struct{})
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	go func() {
		<-sigs
		close(done)
		close(jobs)
		wg.Wait()
		os.Exit(0)
	}()
	defer fmt.Println("You pressed ctrl + C. Shutting down...") // печатаем заключительное информационное сообщение о завершении работы
	//спавним воркеров
	spawnWorkers(n, jobs, &wg)

	//запускаем бесконечный цикл с записью данных в канал jobs
	Loop(n, jobs, done)

}
