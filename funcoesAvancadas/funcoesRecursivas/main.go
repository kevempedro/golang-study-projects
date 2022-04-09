package main

import "fmt"

func main() {
	posicao := uint(10)

	for i := uint(1); i <= posicao; i++ {
		fmt.Println(fibonacci(i))
	}

}

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return uint(posicao)
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
