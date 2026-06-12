package main

import "fmt"

type Pessoa struct {
	Nome          string
	Idade         int
	Nacionalidade string
}

func main() {
	dadosPessoas()
}

func dadosPessoas() {
	carlos := Pessoa{Nome: "Carlos Miguel", Idade: 25, Nacionalidade: "Brasil"}
	fmt.Println(carlos)

	matheus := Pessoa{Nome: "Matheus Silva", Idade: 20, Nacionalidade: "Brasil"}
	fmt.Println(matheus)

	// Utilizando map junto com struct
	pessoas := map[int]Pessoa{
		1: {Nome: "Pedro Nascimento", Idade: 21, Nacionalidade: "Canada"},
		2: {Nome: "Gustavo Lopes", Idade: 26, Nacionalidade: "Japao"},
		3: {Nome: "Antonio Miranda", Idade: 40, Nacionalidade: "Inglaterra"},
	}

	fmt.Println(pessoas)
}
