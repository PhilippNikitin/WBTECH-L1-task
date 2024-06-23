package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Определяем большие числа для a и b
	a := new(big.Int).Exp(big.NewInt(2), big.NewInt(25), nil)
	b := new(big.Int).Exp(big.NewInt(2), big.NewInt(25), nil)

	// Выполняем арифметические операции
	sum := new(big.Int).Add(a, b)
	diff := new(big.Int).Sub(a, b)
	product := new(big.Int).Mul(a, b)
	quotient := new(big.Int).Div(a, b)

	// Выводим результаты
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("a + b =", sum)
	fmt.Println("a - b =", diff)
	fmt.Println("a * b =", product)
	fmt.Println("a / b =", quotient)
}
