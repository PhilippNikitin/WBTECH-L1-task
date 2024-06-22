package main

import "fmt"

// функция binarySearch - реализует бинарный поиск

// Параметры

// slice - отсортированный по возрастанию слайс, x - число, индекс которого необходимо определить
// возвращает индекс искомого элемента
func binarySearch(slice []int, x int) {
	left := 0               // инициализируем начальное значение левой границы поиска
	right := len(slice) - 1 // инициализируем начальное значение правой границы поиска
	var mid int             // объявляем средний элемент
	for {
		mid = (left + right) / 2 // инициализируем средний элемент
		fmt.Println("mid", mid)
		if slice[mid] < x {
			left = mid + 1
			fmt.Println("left", left)
		} else if slice[mid] > x {
			right = mid - 1
			fmt.Println("right", right)
		} else { // slice[mid] == x
			fmt.Println("Индекс загаданного числа в слайсе:", mid)
			break
		}
	}

}

func main() {
	var size int // объявляем переменную, которая будет определять размер слайса
	fmt.Println("Введите размер слайса (натуральное число)")
	fmt.Scanln(&size)
	newSlice := make([]int, size) // объявляем слайс размером 100
	for i := 0; i < size; i++ {   // заполняем слайс числами от 1 до 100 включительно
		newSlice[i] = i + 1
	}
	fmt.Printf("Инициализирован слайс newSlice, содержащий числа от 1 до %d включительно.\n", size)

	var num int // объявляем переменную, в которую будем записывать число введенное пользователем
	fmt.Printf("Введите число от 1 до %d, индекс которого Вы хотите определить:\n", size)
	fmt.Scanln(&num)
	if num > size {
		fmt.Println("Вы ввели число больше заданного диапазона.")
	} else {
		binarySearch(newSlice, num)
	}
}
