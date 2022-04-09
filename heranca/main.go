package main

import "fmt"

func main() {
	p := pessoa{"Kevem", 24}
	fmt.Println(p)

	e := estudante{p, "1234"}
	fmt.Println(e.nome)
}

type pessoa struct {
	nome  string
	idade int
}

type estudante struct {
	pessoa
	matricula string
}
