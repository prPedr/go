package main

import "fmt"

func main() {
	// var -> Podem ter seus valores mudados
	var nome = "Carlos"
	var idade = 30

	fmt.Printf("%s possuiu %d anos de idade. \n", nome, idade)

	// Declarando uma variavel e atribuindo um tipo de dado a ela (So é permitido em var em const nao pode)
	var teste_frase string

	teste_frase = "Teste de escita"

	fmt.Println(teste_frase)

	// const -> Nao podem ter seus valores alterados

	const pi = 3.14
	fmt.Printf("Valor de PI: %.2f", pi)

	const url = "https://appowl.com.br"
	fmt.Println(url)
}
