package main

import "fmt"

func main() {
	s := soma(1, 2, 3)

	fmt.Println(s)

	escrever("Números", 1, 2, 3, 4)
}

func soma(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}
