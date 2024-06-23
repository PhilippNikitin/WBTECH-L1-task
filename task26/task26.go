package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func checkUnique(s string) bool {
	charSet := map[rune]struct{}{}
	for _, char := range s {
		_, found := charSet[char]
		if !found {
			charSet[char] = struct{}{}
		} else {
			return false
		}
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите строку, которую необходимо проверить на предмет уникальности символов и нажмите Enter:")
	input, _ := reader.ReadString('\n')
	fmt.Println(checkUnique(strings.TrimSpace(input)))
}
