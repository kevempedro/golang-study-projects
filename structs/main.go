package main

import "fmt"

func main() {
	var u usuario
	fmt.Println(u)

	u.nome = "Kevem Pedro"
	u.idade = 24

	fmt.Println(u)

	e := endereco{"Rau Oito", "Cohab"}

	u2 := usuario{"Kevem Lima", 24, e}

	fmt.Println(u2)

	u3 := usuario{nome: "Kevem"}
	fmt.Println(u3)
}

type usuario struct {
	nome     string
	idade    int
	endereco endereco
}

type endereco struct {
	rua    string
	bairro string
}
