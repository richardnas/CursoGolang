package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	frutas := []string{"Banana", "Laranja", "Maça", "Manga"}

	resultado, _ := json.Marshal(frutas)
	fmt.Println(string(resultado))
}
