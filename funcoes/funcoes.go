package main

import (
	"errors"
	"fmt"
)

func somarNumeros() float64 {
	var numero1, numero2 float64

	fmt.Print("Informe o primeiro número: ")
	fmt.Scan(&numero1)

	fmt.Print("Informe o segundo número: ")
	fmt.Scan(&numero2)

	soma := numero1 + numero2

	fmt.Printf("%.2f + %.2f = %.2f", numero1, numero2, soma)

	return soma
}

type Usuario struct {
	Nome  string
	Senha string
}

func loginSistema() error {
	usuario1 := Usuario{Nome: "Pedro.Cardoso", Senha: "Pedro@123"}

	var nomeInput string
	var senhaInput string

	fmt.Print("Digite o seu usuario: ")
	fmt.Scan(&nomeInput)

	fmt.Print("Digite sua senha: ")
	fmt.Scan(&senhaInput)

	if usuario1.Nome != nomeInput {
		return errors.New("Usuario invalido")
	}

	if usuario1.Senha != senhaInput {
		return errors.New("Senha invalida")
	}

	return nil
}

func main() {
	if err := loginSistema(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Login realizado")
}
