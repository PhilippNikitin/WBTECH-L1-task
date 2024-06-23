package main

import (
	"fmt"
	"time"
)

// Внутри sleep() функции используем time.After(duration), которая возвращает канал, который будет закрыт по истечении указанной продолжительности.
// Используем оператор <- для ожидания закрытия этого канала.
func sleep(duration time.Duration) {
	<-time.After(duration)
}

func main() {
	fmt.Println("Starting...")
	sleep(2 * time.Second)
	fmt.Println("Done!")
}
