package main

import (
	"fmt"

	"github.com/richardnas/go-examples/16_visibilidade/animal"
)

func main() {
	g := animal.Gato{Nome: "Almôndega"}
	fmt.Println(g.Nome)
}
