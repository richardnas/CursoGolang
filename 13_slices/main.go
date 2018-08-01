package main

import "fmt"

func main() {
	frutas := [4]string{"Banana", "Laranja", "Maça", "Manga"}

	favoritas := frutas[1:3]

	for _, fruta := range favoritas {
		fmt.Println(fruta)
	}
}
