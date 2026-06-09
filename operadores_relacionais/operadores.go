package main

import "fmt"

func main() {
	compararNumeros()
}

func compararNumeros() {
	var numero1 float64
	var numero2 float64

	fmt.Print("Digite um número para comparacao: ")
	fmt.Scan(&numero1)

	fmt.Print("Digite outro numero para comparacao: ")
	fmt.Scan(&numero2)

	if numero1 == numero2 {
		fmt.Printf("O número 1: %.2f é igual ao número 2: %.2f \n", numero1, numero2)
	}

	if numero1 > numero2 {
		fmt.Printf("O número 1: %.2f é maior que o número 2: %.2f \n", numero1, numero2)
	}

	if numero1 < numero2 {
		fmt.Printf("O número 1: %.2f é menor que o número 2: %.2f \n", numero1, numero2)
	}

	if numero1 != numero2 {
		fmt.Printf("O número 1: %.2f é diferente do número 2: %.2f \n", numero1, numero2)
	}
}
