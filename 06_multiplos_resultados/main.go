package main

import (
	"fmt"
)

func palavras() (string, string) {
	return "Hello", "World"
}

func main() {
	h, w := palavras()
	fmt.Println(h, w)
}
