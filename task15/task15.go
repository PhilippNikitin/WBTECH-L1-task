package main

import (
	"bytes"
	"fmt"
	"strings"
)

/*
// исходный код
var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100] // изменяем глобальную переменную внутри функции
}
*/

/*
v := createHugeString(1 << 10)
Функция createHugeString создает большую строку (1 Кб) и возвращает ее.
Функция someFunc сохраняет первые 100 символов этой строки в глобальную переменную justString.
Оставшаяся часть строки, созданной в createHugeString, остается в памяти, но на нее нет больше ссылок, поэтому она не может быть удалена сборщиком мусора. Это приводит к утечке памяти.
Создание и хранение строки длиной 1024 символа в памяти, когда нужно только первые 100 символов, является неоптимальным с точки зрения использования памяти.
*/

/*
Почему плохо использовать глобальные переменные (отдельные причины)

1. В большом проекте при обилии глобальных переменных возникает путаница в именах. Глобальную переменную же видно отовсюду, надо, чтобы отовсюду было понятно, зачем она.
2. Глобальные переменные ухудшают масштабируемость проекта.
3. Глобальные переменные ухудшают читаемость кода (в каком-то конкретно взятом месте непонятно, нужна ли какая-то конкретная глобальная переменная, или нет).
4. Глобальные переменные приводят к трудноуловимым ошибкам. Примеры: нежелательное изменение её значения в другом месте/другим потоком,
ошибочное использование глобальной переменной для промежуточных вычислений из-за совпадения имен, возвращение функцией неправильного значения при тех же параметрах
(оказывается, она зависима от глобальной переменной, а ее кто-то поменял).
5. Глобальные переменные увеличивают число прямых и косвенных связей в системе, делая её поведение труднопредсказуемым, а её саму - сложной для понимания и развития.

Источник: https://ru.stackoverflow.com/questions/510910/%D0%9F%D0%BE%D1%87%D0%B5%D0%BC%D1%83-%D0%B3%D0%BB%D0%BE%D0%B1%D0%B0%D0%BB%D1%8C%D0%BD%D1%8B%D0%B5-%D0%BF%D0%B5%D1%80%D0%B5%D0%BC%D0%B5%D0%BD%D0%BD%D1%8B%D0%B5-%D1%8D%D1%82%D0%BE-%D0%B7%D0%BB%D0%BE-%D0%B0-%D0%BF%D0%BE%D0%BB%D1%8F-%D0%BA%D0%BB%D0%B0%D1%81%D1%81%D0%B0-%D0%BD%D0%B5%D1%82
*/

// Возможная оптимизация

// 1. заранее указываем, строка какой длины нам требуется, через параметр size
func someFunc(size int) string {
	return createHugeString(size)
}

// 2. используем буфер (bytes.Buffer) для оптимизации использования памяти. Метод Grow выделяет необходимую память заранее.
func someFunc2(size int) string {
	var buf bytes.Buffer
	buf.Grow(100)
	buf.WriteString(createHugeString(size))
	return buf.String()
}

func createHugeString(size int) string {
	return strings.Repeat("a", size)
}

func main() {
	var justString string
	justString = someFunc(100) // создаем строку именно того размера, который нам нужен
	fmt.Println(justString)

	var anotherString string
	anotherString = someFunc2(100) // создаем строку именно того размера, который нам нужен
	fmt.Println(anotherString)
}
