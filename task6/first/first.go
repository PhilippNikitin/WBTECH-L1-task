// using channels for signal termination

package first

import (
	"fmt"
	"time"
)

var StopChannel = make(chan bool)

// определяем функцию, которая будет выполняться во время действия горутины

func FirstProcess(stopChannel chan bool) {
	fmt.Println("Начало работы горутины 1.")
	for {
		select {
		case <-stopChannel: // когда из канала stopChannel приходит значение true
			fmt.Println("Получен сигнал о завершении работы горутины 1. Горутина завершается.")
			return // выходим их функции
		default: // сигнала о завершении работы горутины не поступало
			fmt.Println("Горутина 1 продолжает работать")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
