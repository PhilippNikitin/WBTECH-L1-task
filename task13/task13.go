package main

import "fmt"

// first approach - multi-value assignment
func interchangeMultiValAssignment(x, y int) (int, int) {
	x, y = y, x
	return x, y
}

// second approach - use of addition and subtraction
func interchangeAddSub(x, y int) (int, int) {
	y = x + y
	x = y - x
	y = y - x
	return x, y
}

func main() {
	x, y := 2, 10 // инициализируем переменные с числами
	fmt.Println("Значения переменных до обмена:")
	fmt.Printf("x = %d\ny = %d\n", x, y) // выводим начальные значения x и y в stdout

	x, y = interchangeMultiValAssignment(x, y)
	fmt.Println("Значения переменных после первого обмена:")
	fmt.Printf("x = %d\ny = %d\n", x, y) // выводим значения x и y после обмена в stdout

	x, y = interchangeAddSub(x, y)
	fmt.Println("Значения переменных после второго обмена:")
	fmt.Printf("x = %d\ny = %d\n", x, y) // выводим значения x и y после обмена в stdout
}
