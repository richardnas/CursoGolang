package main

import (
	"fmt"
)

func main() {
	animal := "gato"

	switch animal {
	case "gato":
		fmt.Println("Esse animal é um gatinho!")
	case "cachorro":
		fmt.Println("Esse animal é um cachorrinho!")
	default:
		fmt.Println("Não sei qual é esse animal! :(")
	}
}
