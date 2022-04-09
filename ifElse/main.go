package main

import "fmt"

func main() {
	numero := 10

	if numero > 15 {
		fmt.Println("Maior")
	} else {
		fmt.Println("Menor")
	}

	if numero2 := numero; numero2 > 0 {
		fmt.Println("Maior")
	} else if numero2 < 10 {
		fmt.Println("Menor")
	} else {
		fmt.Println("Nada")
	}
}
