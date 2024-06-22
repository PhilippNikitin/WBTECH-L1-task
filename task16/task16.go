package main

import "fmt"

func quicksort(unsorted []int) []int {
	if len(unsorted) <= 1 { // если слайс пустой или содержит один элемент - возвращаем данный слайс без изменений
		return unsorted
	} else if len(unsorted) == 2 {
		if unsorted[0] > unsorted[1] { // если элемент с индексом 0 меньше, или равен элементу с индексом 1
			unsorted[0], unsorted[1] = unsorted[1], unsorted[0] // обмениваем значения переменных
		}
		return unsorted
	} // если слайс состоит из трех элементов или больше
	pivot := unsorted[0]      // создаем опорный элемент, в качестве которого выбираем первый элемент слайса
	less := []int{}           // создаем слайс, в котором будем хранить элементы меньше опорного
	equalOrGreater := []int{} // создаем слайс, в котором будем хранить элементы больше опорного
	for _, val := range unsorted[1:] {
		if val < pivot { // если рассматриваемый элемент меньше опорного, необходимо перенести его в начало слайса
			less = append(less, val)
		} else { // if val >= support
			equalOrGreater = append(equalOrGreater, val)
		}
	}
	return append(append(quicksort(less), []int{pivot}...), quicksort(equalOrGreater)...)
}

func main() {
	unsorted := []int{8, 7, 19, 0, -3, 547, -900, 2345, -9, 0, 5} // попробовать работу функции с пустым слайсом или со слайсом, состоящим из одного элемента
	sorted := quicksort(unsorted)
	fmt.Println(sorted)
}
