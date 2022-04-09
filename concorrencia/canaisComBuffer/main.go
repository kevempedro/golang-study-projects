package main

import "fmt"

func main() {
	canal := make(chan string, 2)

	canal <- "Hellow World!"
	canal <- "Programando em Go!"

	mensagem := <-canal
	mesagem2 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mesagem2)
}
