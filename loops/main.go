package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 10 {
		fmt.Println(i)
		i++
		time.Sleep(time.Second)
	}

	for j := 0; j < 10; j++ {
		fmt.Println(j)
		time.Sleep(time.Second)
	}

	nomes := [3]string{"Kevem", "Lima", "Pedro"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for indice, letra := range "palavra" {
		fmt.Println(indice, string(letra))
	}

	usuarios := map[string]string{
		"nome": "Kevem",
		"lima": "Lima",
	}

	for chave, valor := range usuarios {
		fmt.Println(chave, valor)
	}

	for {
		fmt.Println("Execuntando infinitamente")
		time.Sleep(time.Second)
	}

	fmt.Println("Execuntando infinitamente")
}
