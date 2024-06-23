package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите строку, которую необходимо перевернуть, и нажмите Enter:")
	input, _ := reader.ReadString('\n')

	fmt.Println(reverse(input))
}
