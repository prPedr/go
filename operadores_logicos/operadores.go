package main

import "fmt"

func main() {
	estoque := true
	vendaLiberada := false

	if estoque && vendaLiberada {
		fmt.Println("Produto em separacao.")
	}

	if estoque == false {
		fmt.Println("O produto nao pode ser colocado em separacao pois nao tem estoque.")
	}

	if estoque && vendaLiberada == false {
		fmt.Println("O produro nao pode ser colocado em separacao pois a venda nao foi liberada.")
	}
}
