package main

import "fmt"

func main() {
	fmt.Println(verificaAlunoAprovado(6, 6))
	fmt.Println("Pós execução!")
}

func verificaAlunoAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A média é exatamente 6!")
}

func recuperarExecucao() {
	fmt.Println("Tentativa de recuperar a execução")

	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}
