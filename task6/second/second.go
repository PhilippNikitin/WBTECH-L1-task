// use shared variables for goroutine termination

package second

import (
	"fmt"
	"sync"
	"time"
)

var StopIt bool = false
var WG sync.WaitGroup

func SecondProcess(stopIt *bool, wg *sync.WaitGroup) {
	fmt.Println("Начало работы горутины 2.")
	defer wg.Done()

	for {
		if *stopIt {
			fmt.Println("Получен сигнал о завершении горутины 2. Завершаем выполнение горутины 2 через механизм общих переменных.")
			return
		}

		fmt.Println("Горутина 2 продолжает работать.")
		time.Sleep(500 * time.Millisecond)
	}
}
