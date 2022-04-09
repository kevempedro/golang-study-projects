package main

import "fmt"

func main() {
	soma := somar(1, 2)
	fmt.Println(soma)

	var fun = func(nome string) {
		fmt.Printf("%v\n", nome)
	}

	fun("Kevem")

	resultadoSoma, resultadoDivisao := calculos(10, 4)

	fmt.Println(resultadoSoma, resultadoDivisao)
}

func somar(n1 int, n2 int) int {
	return n1 + n2
}

func calculos(n1, n2 int) (int, float64) {
	soma := n1 + n2
	divisao := n1 / n2

	return soma, float64(divisao)
}
