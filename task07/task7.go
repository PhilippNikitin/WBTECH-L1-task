package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создаём map и mutex
	myMap := make(map[string]int)
	var mutex sync.Mutex

	// Создаём горутины для записи данных в map
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// Блокируем доступ к map
			mutex.Lock()
			defer mutex.Unlock()

			// Записываем данные в map
			myMap[fmt.Sprintf("key_%d", id)] = id
			fmt.Printf("Горутина %d записала данные в map\n", id)
		}(i)
	}

	// Ждём завершения всех горутин
	wg.Wait()

	// Выводим содержимое map
	fmt.Println("Содержимое map:")
	for key, value := range myMap {
		fmt.Printf("%s: %d\n", key, value)
	}
}
