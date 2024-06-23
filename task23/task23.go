package main

import "fmt"

func delElem(slice []int, i int) []int {
	if len(slice) == 0 {
		fmt.Println("Длина слайса равна 0. Возвращаю неизменный слайс.")
		return slice
	}
	return append(slice[:i], slice[i+1:]...)
}

func main() {
	a := []int{1, 2, 3}
	b := []int{1}
	c := []int{}

	fmt.Println(delElem(a, 2))
	fmt.Println(delElem(b, 0))
	fmt.Println(delElem(c, 0))
}
