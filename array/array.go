package main

import "fmt"

func main() {
	array1()
	array2()
	array3()
}

func array1() {
	var numeros [4]int

	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30
	numeros[3] = 40

	fmt.Println(numeros)
}

func array2() {
	numeros := [4]int{10, 20, 30, 40}

	fmt.Println(numeros)
}

func array3() {
	permissoes := [3]string{"usuario", "admin", "editor"}

	fmt.Print(permissoes)
}
