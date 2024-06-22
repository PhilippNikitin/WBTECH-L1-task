package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var s struct{}
	fmt.Println(unsafe.Sizeof(s))
	var i interface{}
	fmt.Println(unsafe.Sizeof(i))
	var b bool
	fmt.Println(unsafe.Sizeof(b))
	var str string
	fmt.Println(unsafe.Sizeof(str))
}
