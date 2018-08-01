package main

import (
	"fmt"
)

type endereco struct {
	logradouro, bairro, cidade, estado, cep, numero string
}

type pessoa struct {
	endereco
}

func main() {
	p := pessoa{
		endereco{
			logradouro: "Rua XYZ",
			bairro:     "Abcdefgh",
			cidade:     "Lins",
			estado:     "São Paulo",
			cep:        "16400-000",
			numero:     "100",
		},
	}

	fmt.Println(p.cep)
}
