package main

import "fmt"

func main() {
	usuario := usuario{"Kevem", 24}
	usuario.salvar()

	fmt.Println("Usuário é maior de idade: ", usuario.verificaSeEMaiorDeIdade())

	usuario.fazerAniversario()

	fmt.Println(usuario.idade)
}

type usuario struct {
	nome  string
	idade int
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuário %s no banco de dados\n", u.nome)
}

func (u usuario) verificaSeEMaiorDeIdade() bool {
	if u.idade > 18 {
		return true
	}

	return false
}

func (u *usuario) fazerAniversario() {
	u.idade++
}
