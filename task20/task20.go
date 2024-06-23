package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverseWords(s string) string {
	wordSlice := strings.Split(s, " ")
	reverseWords := []string{}
	for i := len(wordSlice) - 1; i > -1; i-- {
		reverseWords = append(reverseWords, wordSlice[i])
	}
	return strings.Join(reverseWords, " ")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите строку, слова в которой необходимо перевернуть, и нажмите Enter:")
	input, _ := reader.ReadString('\n')

	fmt.Println(reverseWords(strings.TrimSpace(input)))
}
