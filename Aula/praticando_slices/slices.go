package main

import "fmt"

func main() {
	var tarefas []string

	// Adicionando itens no slice
	tarefas = append(tarefas, "Estudar GO", "Levar o cachorro para passear", "Ir no mercado", "Ir na academia")
	fmt.Println(tarefas)

	// Mostrando o tamanho do slice
	fmt.Println(len(tarefas))

	// Removendo o primeiro item
	tarefas = tarefas[1:]
	fmt.Println(tarefas)

	// Removendo o ultimo item
	tarefas = tarefas[:len(tarefas)-1]
	fmt.Println(tarefas)

	// Removendo o item do meio
	tarefas = append(tarefas[:1], tarefas[2:]...)
	fmt.Println(tarefas)
}
