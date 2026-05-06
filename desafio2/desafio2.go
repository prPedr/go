package main

import "fmt"

func main() {
	var p float64
	var i float64
	var n float64

	fmt.Print("Digite o valor capital: ")
	fmt.Scan(&p)

	fmt.Print("")

	fmt.Print("Digite a taxa de juros: ")
	fmt.Scan(&i)

	calcularPorcentagem := i / 100

	i = calcularPorcentagem

	fmt.Print("")

	fmt.Print("Digite o numero de meses: ")
	fmt.Scan(&n)

	j := p * i * n

	fmt.Printf("Total de juros: %.2f", j)
}
