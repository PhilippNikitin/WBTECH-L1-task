package main // Struct embedding

import "fmt"

type Human struct { // определяем структуру Human
	X float64 // задаем поле X для Human
	Y float64 // задаем поле Y для Human
}

func (h *Human) Move(x, y float64) { // определяем функцию Move, которая будет принимать два аргумента X и Y и прибавлять их к существующим координатам
	h.X += x
	h.Y += y
}

func (h *Human) ReturnToTheBase() { // определяем функцию для обнуления координат
	h.X = 0
	h.Y = 0
}

type Action struct { // определяем структуру Action
	*Human // устанавливаем указатель на структуру Human для реализации embedding
}

func NewAction() *Action { // определяем функцию для создания нового экземпляра Action
	return &Action{
		Human: &Human{},
	}
}

func main() { // проверяем в действии механизм внедрения методов
	action := NewAction()

	action.Move(9.6, 10.8)
	fmt.Println(action.Human)

	action.Move(100., 100.)
	fmt.Println(action.Human)

	action.ReturnToTheBase()
	fmt.Println(action.Human)
}
