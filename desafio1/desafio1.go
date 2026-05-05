package main

import "fmt"

func main() {
	calcularIdade()
	converterHorasParaMinutos()
}

func calcularIdade() {
	var anoAtual int
	var anoNascimento int

	fmt.Print("Digite o ano atual: ")
	fmt.Scan(&anoAtual)

	fmt.Print("Digite seu ano de nascimento: ")
	fmt.Scan(&anoNascimento)

	resultadoIdade := anoAtual - anoNascimento

	fmt.Printf("Sua idade é: %d", resultadoIdade)
}

func converterHorasParaMinutos() {
	var horas float64

	fmt.Printf("Digite a hora: ")
	fmt.Scan(&horas)

	conversaoMinutos := horas * 60

	fmt.Printf("%.2f é igual há %.2f minutos", horas, conversaoMinutos)
}
