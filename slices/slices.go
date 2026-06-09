package main

import "fmt"

func main() {
	var lista []int

	lista = append(lista, 10, 20, 30, 40, 50)

	fmt.Println(lista)

	// Criacao do Slices de frutas.
	frutas := []string{"maca, banana, mamao"}
	fmt.Println(frutas)

	// Slices de fruta apos o append
	frutas = append(frutas, "uva, morango")
	fmt.Println(frutas)
}
