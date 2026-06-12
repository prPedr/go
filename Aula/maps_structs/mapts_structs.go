package main

import "fmt"

type Usuario struct {
	Nome      string
	Permissao string
	Status    bool
	Setor     string
}

func main() {
	dadosUsuarios()
}

func dadosUsuarios() {
	usuarios := map[int]Usuario{
		1: {Nome: "Pedro Nascimento", Permissao: "Admin", Status: true, Setor: "TI"},
		2: {Nome: "Carlos Chagas", Permissao: "Admin", Status: false, Setor: "TI"},
		3: {Nome: "Ana Souza", Permissao: "Editor", Status: true, Setor: "Logistica"},
		4: {Nome: "Antonio Silva", Permissao: "User", Status: true, Setor: "Compras"},
		5: {Nome: "Mariana Miranda", Permissao: "User", Status: false, Setor: "Enfermangem"},
		6: {Nome: "Gabriela Leme", Permissao: "Editor", Status: true, Setor: "Marketing"},
	}

	// Adicionando um novo item em uma struct
	usuarios[len(usuarios)+1] = Usuario{Nome: "Henrique Lopes", Permissao: "Admin", Status: true, Setor: "TI"}

	i := 1

	for i <= len(usuarios) {
		fmt.Printf("Nome: %s | Permissao: %s | Status: %t | Setor: %s \n", usuarios[i].Nome, usuarios[i].Permissao, usuarios[i].Status, usuarios[i].Setor)
		i++
	}
}
