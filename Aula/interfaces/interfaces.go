package main

import "fmt"

type Animal interface {
	EmitirSom()
}

func (c Cachorro) EmitirSom() {
	fmt.Println("Teste1")
}

func (a Gato) EmitirSom() {
	fmt.Println("Teste2")
}

type Cachorro struct{}

type Gato struct{}

func fazerEmitirSom(a Animal) {
	a.EmitirSom()
}

func main() {
	cachorro1 := Cachorro{}
	gato1 := Gato{}

	fazerEmitirSom(cachorro1)
	fazerEmitirSom(gato1)
}
