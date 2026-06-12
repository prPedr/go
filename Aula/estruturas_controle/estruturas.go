package main

import "fmt"

func main() {
	var nota1 float64
	var nota2 float64
	var nota3 float64

	fmt.Print("Digite a primeira nota: ")
	fmt.Scan(&nota1)

	fmt.Print("Digite a segunda nota: ")
	fmt.Scan(&nota2)

	fmt.Print("Digite a terceira nota: ")
	fmt.Scan(&nota3)

	media := (nota1 + nota2 + nota3) / 3

	if media >= 7 {
		fmt.Println("APROVADO!")
		fmt.Printf("Media: %.2f", media)
	}

	if media < 7 && media >= 5 {
		fmt.Println("RECUPERACAO")
		fmt.Printf("Media: %.2f", media)
	}

	if media < 5 {
		fmt.Println("PROVADO")
		fmt.Printf("Media: %.2f", media)
	}
}
