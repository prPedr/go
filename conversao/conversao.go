package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Converter: Int -> Float
	var valor1 = 30
	var valor1Convertido1 = float64(valor1)
	fmt.Printf("Valor 1 = %.2f \n", valor1Convertido1)

	// Converter: int -> String
	var valor1Convertido2 = strconv.Itoa(valor1)
	fmt.Printf("Valor 2 = %s \n", valor1Convertido2)

	// Converter: String -> Int
	valor2 := "80"
	valorInteiro, err := strconv.Atoi(valor2)

	if err != nil {
		fmt.Println("Erro ao realizar conversao")
	} else {
		fmt.Println(valorInteiro)
	}

	// Converter: String -> Float
	pi := "3.14"
	piConvertido, err2 := strconv.ParseFloat(pi, 64)

	if err2 != nil {
		fmt.Println("Erro ao realizar conversao")
	} else {
		fmt.Println(piConvertido)
	}
}
