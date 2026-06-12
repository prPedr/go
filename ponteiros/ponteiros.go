package main

import "fmt"

func main() {
	numero := 60

	var p *int = &numero

	fmt.Println("Valor do ponteiro na memoria P: ", p)
	fmt.Println("Valor apontado por P: ", *p)
}
