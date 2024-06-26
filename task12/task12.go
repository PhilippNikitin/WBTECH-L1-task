package main

import "fmt"

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"} // объявляем и инициализируем слайс с заданными словами
	new_set := map[string]struct{}{}                      // объявляем и инициализируем множество для сохранения уникальных слов

	for _, word := range words {
		_, found := new_set[word]
		if !found {
			new_set[word] = struct{}{}
		}
	}

	fmt.Println(new_set)

}
