package main

import "fmt"

func main() {
	generica("kevem")
	generica(10)
	generica(true)
}

func generica(interf interface{}) {
	fmt.Println(interf)
}
