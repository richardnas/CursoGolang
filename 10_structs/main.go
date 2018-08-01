package main

import "fmt"

type Cachorro struct {
	Nome  string
	Idade int
}

func main() {
	c := Cachorro{"Rex", 5}
	fmt.Println(c.Nome)
	fmt.Println(c.Idade)

}
