package main

import "fmt"

// класс собака
type Dog struct{}

// реакция собаки
func (dog *Dog) Woof() {
	fmt.Println("Гав-гав.")
}

// класс кошка
type Cat struct{}

// реакция кошки - издает звуки, только если ее позвать
func (cat *Cat) Meow(isCall bool) {
	if isCall {
		fmt.Println("Мяу-мяу.")
	}
}

// целевой интерфейс - Target
type AnimalAdapter interface {
	Reaction() // общая сигнатура вызова - Reaction()
}

// адаптер для собаки
type DogAdapter struct {
	*Dog
}

// реакция собаки
func (adapter *DogAdapter) Reaction() {
	adapter.Woof()
}

// конструктор адаптера для собаки, возвращает тип целевого интерфейса - AnimalAdapter
func NewDogAdapter(dog *Dog) AnimalAdapter {
	return &DogAdapter{dog}
}

// адаптер для кошки
type CatAdapter struct {
	*Cat
}

// реакция кошки
func (adapter *CatAdapter) Reaction() {
	// адаптер автоматически зовет кошку isCall = true
	adapter.Meow(true)
}

// конструктор адаптера для кота, возвращает тип целевого интерфейса - AnimalAdapter
func NewCatAdapter(cat *Cat) AnimalAdapter {
	return &CatAdapter{cat}
}

// класс - ворона
type Crow struct {
}

// реакция вороны - адаптер не нужен, нужный метод итак есть
func (c *Crow) Reaction() {
	fmt.Println("Кар-кар.")
}

func main() {
	// инициализируем переменную animals - массив из трех элементов типа AnimalAdapter
	animals := [3]AnimalAdapter{NewDogAdapter(&Dog{}), NewCatAdapter(&Cat{}), &Crow{}}

	// слушаем, что говорит каждое животное
	for _, animal := range animals {
		animal.Reaction()
	}
}
