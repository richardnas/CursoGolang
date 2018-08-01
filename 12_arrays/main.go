package main

import "fmt"

func main() {
	frutas := [4]string{"Banana", "Laranja", "Maça", "Manga"}
	for _, fruta := range frutas {
		fmt.Println(fruta)
	}
}
