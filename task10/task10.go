package main

import "fmt"

func defineGroup(t float32) int {
	return (int(t) / 10) * 10

}

func main() {
	tempSlice := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	tempGroups := map[int][]float32{} // инициализируем слайс с группой и соответсвующими температурами

	for _, t := range tempSlice { // итерируемся по слайсу с температурами
		group := defineGroup(t) // определяем группу, которая соответствует каждой температуре
		// добавляем в словарь значение для соответствующей группы;
		// если ключа нет в словаре, он будет создан, а значение - добавлено к пустому слайсу []float32{}
		// если ключ есть, то для него обновится значение - добавится значение t к соответствующему слайсу
		tempGroups[group] = append(tempGroups[group], t)
	}

	// выводим результаты
	for group, temps := range tempGroups {
		fmt.Printf("%d: %v\n", group, temps)
	}
}
