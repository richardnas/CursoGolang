package main

import "fmt"

func maisUm(num *int) {
	*num = *num + 1
}

func main() {
	valor := 17

	fmt.Println(valor)

	maisUm(&valor)

	fmt.Println(valor)
}
