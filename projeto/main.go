package main

import (
	"fmt"
	"projeto/aluno"
)

func main() {
	aluno1 := aluno.ObterInformacoesAluno()

	fmt.Println(aluno1.RetornoInformacoesAluno())
	fmt.Println(aluno1.Status())
}
