package main

import (
	"fmt"
	"task6/first"
	"task6/fourth"
	"task6/second"
	"task6/third"
	"time"
)

func main() {
	// иллюстрация завершения работы горутины через механизм каналов
	go first.FirstProcess(first.StopChannel)
	time.Sleep(2 * time.Second)
	first.StopChannel <- true
	fmt.Println("Завершение работы горутины через механизм каналов.")

	// иллюстрация завершения работы горутины через механизм общих переменных
	second.WG.Add(1)
	go second.SecondProcess(&second.StopIt, &second.WG)
	time.Sleep(2 * time.Second)
	second.StopIt = true
	second.WG.Wait()

	// иллюстрация механизма завершения работы горутины при помощи закрытия канала
	go third.ThirdProcess(third.StopChannel)
	time.Sleep(2 * time.Second)
	close(third.StopChannel)
	fmt.Println("Заканчиваем выполнение горутины 3, используя закрытие канала как сигнал завершения работы.")

	// использование контекста для завершения горутины
	go fourth.FourthProcess(fourth.Ctx)
	time.Sleep(2 * time.Second)
	fourth.Cancel()
	fmt.Println("Заканчиваем выполнение горутины 4 при помощи контекста.")

	fmt.Println("Завершение мейн горутины.")
}
