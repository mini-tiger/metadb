package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := []byte("abcdefghijklmnopqrstuvwsyz")
	b := *(*string)(unsafe.Pointer(&a))
	fmt.Println(b)
	fmt.Printf("%p,%p\n", &a, &b)

	c := *(*[]byte)(unsafe.Pointer(&b))
	fmt.Printf("%p,%p,%p\n", &a, &b, &c)
}
