package main

import "fmt"

// var -> Podem ter seus valores mudados

func main() {
	var nome = "Carlos"
	var idade = 30

	fmt.Printf("%s possuiu %d anos de idade. \n", nome, idade)

	var teste_frase string // Declarando uma variavel e atribuindo um valor a ela

	teste_frase = "Teste de escita"

	fmt.Println(teste_frase)
}
