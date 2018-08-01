package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float32
	perimetro() float32
}

type retangulo struct {
	altura, largura float32
}

func (r retangulo) area() float32 {
	return r.altura * r.largura
}

func (r retangulo) perimetro() float32 {
	return 2*r.altura + 2*r.largura
}

type circulo struct {
	raio float32
}

func (c circulo) area() float32 {
	return math.Pi * c.raio * c.raio
}

func (c circulo) perimetro() float32 {
	return 2 * math.Pi * c.raio
}

func imprimirArea(f forma) {
	fmt.Println("Área =>", f.area())
}

func imprimirPerimetro(f forma) {
	fmt.Println("Perímetro =>", f.perimetro())
}

func main() {
	fmt.Println("Retângulo")

	r := retangulo{altura: 5, largura: 10}
	imprimirArea(r)
	imprimirPerimetro(r)

	fmt.Println("Círculo")
	c := circulo{raio: 5}
	imprimirArea(c)
	imprimirPerimetro(c)
}
