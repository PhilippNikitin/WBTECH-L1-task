package main

import "fmt"

func main() {
	set1 := map[string]string{} // объявляем первое множество
	set2 := map[string]string{} // объявляем второе множество

	for _, ch := range "xabcde" { // заполняем множество 1
		set1[string(ch)] = ""
	}

	for _, ch := range "cdefghx" { // заполняем множество 2
		set2[string(ch)] = ""
	}

	intersection := map[string]string{} // создаем мапу для сохранения пересечения данных двух множеств

	for ch := range set1 { // для каждого ключа в множестве 1
		_, found := set2[ch] // если он найден в множестве 2
		if found {
			intersection[ch] = "" // добавляем в пересечение
		}
	}

	fmt.Println(intersection)
}
