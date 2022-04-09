package main

import "fmt"

func main() {
	// defer funcao() // defer - adiar
	// funcao2()

	fmt.Println(verificaAlunoAprovado(10, 5))
}

func funcao() {
	fmt.Println("Executando função 1")
}

func funcao2() {
	fmt.Println("Executando função 2")
}

func verificaAlunoAprovado(n1, n2 float64) bool {
	defer fmt.Println("Média calculada. Resultado será retornado!")
	fmt.Println("Verificando se o aluno foi aprovado...")

	if (n1+n2)/2 >= 6 {
		return true
	}

	return false
}
