package main

import "fmt"

func main() {
	var variavel int = 10
	var variavel2 int = variavel

	fmt.Println(variavel, variavel2)

	variavel++
	fmt.Println(variavel, variavel2)

	// Ponteiro é uma referência de memória
	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println(variavel3, *ponteiro) // * antes do ponteiro é a desreferenciação

	variavel3++

	fmt.Println(variavel3, *ponteiro)
}
