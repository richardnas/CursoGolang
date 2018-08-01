package main

import (
	"fmt"
)

func main() {
	num := 10
	if num > 0 {
		fmt.Println("Esse número é positivo!")
	} else if num < 0 {
		fmt.Println("Esse número é negativo!")
	} else {
		fmt.Println("Esse número é zero!")
	}
}
