package main

import "fmt"

func main() {
	numero := 20

	fmt.Println(inverterSinal(numero))
	fmt.Println(numero)

	numero2 := 40
	fmt.Println(numero2)
	inverterSinalPorPonteiro(&numero2)
	fmt.Println(numero2)
}

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalPorPonteiro(numero *int) {
	*numero = *numero * -1
}
