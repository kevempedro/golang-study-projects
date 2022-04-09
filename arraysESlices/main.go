package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Array
	var array [5]string
	array[0] = "Kevem"
	fmt.Println(array)

	array2 := [2]string{"Kevem", "Pedro"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4}
	fmt.Println(array3)
	fmt.Println("-----------------------")

	// Slice
	slice := []int{10, 20, 30}

	slice = append(slice, 18)
	fmt.Println((slice))

	slice2 := array3[1:4]
	fmt.Println((slice2))

	fmt.Println(reflect.TypeOf(array3))
	fmt.Println(reflect.TypeOf(slice))

	// Arrays Internos
	fmt.Println("-----------------------")
	slice3 := make([]float64, 10, 11)

	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3)) // cap - capacidade

	slice3 = append(slice3, 1.5)
	slice3 = append(slice3, 2)

	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3)) // cap - capacidade

	slice4 := make([]float64, 5)
	slice4 = append(slice4, 1.5)

	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4)) // cap - capacidade
}
