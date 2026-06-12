package main

import "fmt"

func main() {
	for {
		fmt.Println("|--------------------------------|")
		fmt.Println("|        Escolha uma opcao       |")
		fmt.Println("|--------------------------------|")
		fmt.Println("| 1 - Calcular idade             |")
		fmt.Println("| 2 - Converter horas em minutos |")
		fmt.Println("| 3 - Sair                       |")
		fmt.Println("|--------------------------------|")
		fmt.Println()

		var opcao string
		fmt.Print("Escolha uma opcao: ")
		fmt.Scan(&opcao)

		switch opcao {
		case "1":
			calcularIdade()

		case "2":
			converterHorasParaMinutos()

		case "3":
			fmt.Println("Finalizando o programa")
			return

		default:
			fmt.Println("Opcao invalida.")
		}
	}

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
