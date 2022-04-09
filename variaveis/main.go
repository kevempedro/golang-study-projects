package main

import "fmt"

func main() {
	var a = "Kevem1"
	const b = "Kevem2"
	c := "Kevem3"
	d, e := 1, 3

	fmt.Println(a, b, c, d, e)

	d, e = e, d

	fmt.Println(a, b, c, d, e)
}
