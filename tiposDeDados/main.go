package main

import (
	"errors"
	"fmt"
)

func main() {
	// Inteiros
	numero := 10000000000000
	fmt.Println(numero)

	var numero2 uint = 10
	fmt.Println(numero2)

	/*
		ALIAS

		int32 = rune
		int8 = byte
	*/
	var numero3 rune = 11323334
	fmt.Println(numero3)

	// Floats
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal float64 = 123.45
	fmt.Println(numeroReal)

	// Strings
	var texto string = "Texto"
	fmt.Println(texto)

	char := 'B'
	fmt.Println(char)

	// Boolean
	boleano := true
	fmt.Println(boleano)

	// Error
	var erro error = errors.New("Erro interno no servidor!!!")
	fmt.Println(erro)

	// Valor 0
	var valorZeroString string
	var valorZeroInt int
	var valorZeroFloat float64
	var valorZeroBooleano bool
	var valorZeroError bool

	fmt.Println(valorZeroString)
	fmt.Println(valorZeroInt)
	fmt.Println(valorZeroFloat)
	fmt.Println(valorZeroBooleano)
	fmt.Println(valorZeroError)
}
