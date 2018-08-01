package main

import "fmt"

func main() {
	var cor = map[string]string{
		"Banana":  "Amarelo",
		"Laranja": "Alaranjado",
		"Maça":    "Vermelho",
		"Manga":   "Vermelho",
	}
	fmt.Println(cor["Banana"])
}
