package main

import "fmt"

func main() {
	soma, sub := calculos(5, 2)

	fmt.Println(soma, sub)
}

func calculos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}
