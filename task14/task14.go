package main

import (
	"fmt"
	"reflect"
)

func findType(a interface{}) {

	// Using type switch
	switch a.(type) {
	case int:
		fmt.Println("Type: int, Value:", a.(int))
	case float64:
		fmt.Println("Type: float64, Value:", a.(float64))
	case string:
		fmt.Println("Type: string, Value:", a.(string))
	case bool:
		fmt.Println("Type: bool, Value:", a.(bool))
	case chan int:
		fmt.Println("Type: chan int, Value:", a.(chan int))
	default:
		fmt.Println("Type not found:", a)
	}
}

func main() {

	// первый способ - через определенную в пакете main функцию findType
	values := []interface{}{"Hello", 78, make(chan int), true, 89.8}

	fmt.Println("Определение типа при помощи функции findType, определенной в пакете main")
	for _, val := range values {
		findType(val)
	}
	fmt.Println()

	// второй способ - при помощи функции reflect.TypeOf()
	fmt.Println("Определение типа при помощи функции reflect.TypeOf()")
	for _, val := range values {
		fmt.Printf("Type: %s, Value: %v \n", reflect.TypeOf(val), val)
	}
	fmt.Println()

	// третий способ - использование спецификатора %T в fmt.Printf()
	fmt.Println("Определение типа при помощи функции спецификатора %T")
	for _, val := range values {
		fmt.Printf("Type: %T, Value: %v \n", val, val)
	}
	fmt.Println()

	// четвертый способ - использование функции reflect.ValueOf(x).Kind()
	fmt.Println("Определение типа при помощи reflect.ValueOf().Kind()")
	for _, val := range values {
		fmt.Printf("Type: %s, Value: %v \n", reflect.ValueOf(val).Kind(), val)
	}

}
