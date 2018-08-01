package main

import "fmt"

type retangulo struct {
	altura, largura float32
}

func (r *retangulo) area() float32 {
	return r.altura * r.largura
}

func (r *retangulo) perimetro() float32 {
	return 2*r.altura + 2*r.largura
}

func main() {
	r := retangulo{altura: 5, largura: 10}
	fmt.Println("Área =>", r.area())
	fmt.Println("Perímetro =>", r.perimetro())
}
