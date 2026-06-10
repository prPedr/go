package main

import "fmt"

func calcularMedia(nota1, nota2 float64) float64 {
	return (nota1 + nota2) / 2
}

func obterStatus(media float64) string {
	switch {
	case media >= 7:
		return "APROVADO"
	case media >= 5:
		return "RECUPERACAO"
	default:
		return "REPROVADO"
	}
}

func main() {
	var nomeAluno string
	var nota1, nota2 float64

	fmt.Print("Informe o nome do aluno: ")
	fmt.Scan(&nomeAluno)

	fmt.Print("Informe a primeira nota do aluno: ")
	fmt.Scan(&nota1)

	fmt.Print("Informe a segunda nota do aluno: ")
	fmt.Scan(&nota2)

	media := calcularMedia(nota1, nota2)
	status := obterStatus(media)

	fmt.Printf(
		"Nome aluno: %s\nStatus: %s\nMédia: %.2f\n",
		nomeAluno,
		status,
		media,
	)
}
