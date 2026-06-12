package main

import "fmt"

func declarandoUmPonteiro() {
	numero := 60

	var p *int = &numero

	fmt.Println("Valor do ponteiro na memoria P: ", p)
	fmt.Println("Valor apontado por P: ", *p)
}

func somarValores(num *int) {
	*num = 16
}

func alterandoValoresPonteiros() {
	numero := 10
	fmt.Println("Valor inicial da variavel: ", numero)

	somarValores(&numero)

	fmt.Println("Valor atual da variavel numero: ", numero)
}

func main() {
	alterandoValoresPonteiros()
}
