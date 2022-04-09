package main

import "fmt"

func main() {
	soma := 1 + 2
	subtracao := 2 - 1
	divisao := 10 / 2
	multiplicacao := 10 * 2
	restoDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDivisao)

	// Atribuição '=' ou ':='
	a := 10
	fmt.Println(a)

	// Operadores relacionais >, >= <, <=, == ou !=
	fmt.Println(10 > 9)

	// Operadores lógicos &&, || ou !
	fmt.Println((10 > 9) && (5 < 10))

	// Operadores unários
	numero := 10
	numero++
	numero += 10
	numero *= 2

	fmt.Println(numero)
}
