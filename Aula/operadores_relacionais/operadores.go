package main

import "fmt"

func main() {
	for {
		fmt.Println("|------------------------------------|")
		fmt.Println("|          Escolha uma opcao         |")
		fmt.Println("|------------------------------------|")
		fmt.Println("| 1 - Comparacao de Números          |")
		fmt.Println("| 2 - Comparacao de tamanho de nomes |")
		fmt.Println("| 3 - Comparacao de nomes            |")
		fmt.Println("| 4 - Sair                           |")
		fmt.Println("|------------------------------------|")
		fmt.Println()

		var opcao string
		fmt.Print("Escolha uma opcao: ")
		fmt.Scan(&opcao)

		switch opcao {
		case "1":
			compararNumeros()

		case "2":
			compararTamanhoNomes()

		case "3":
			compararNomes()

		case "4":
			fmt.Println("Finalizando o programa")
			return

		default:
			fmt.Println("Opcao invalida.")
		}
	}
}

func compararNumeros() {
	var numero1 float64
	var numero2 float64

	fmt.Print("Digite um número para comparacao: ")
	fmt.Scan(&numero1)

	fmt.Print("Digite outro numero para comparacao: ")
	fmt.Scan(&numero2)

	if numero1 == numero2 {
		fmt.Printf("O número 1: %.2f é igual ao número 2: %.2f \n", numero1, numero2)
	}

	if numero1 > numero2 {
		fmt.Printf("O número 1: %.2f é maior que o número 2: %.2f \n", numero1, numero2)
	}

	if numero1 < numero2 {
		fmt.Printf("O número 1: %.2f é menor que o número 2: %.2f \n", numero1, numero2)
	}

	if numero1 != numero2 {
		fmt.Printf("O número 1: %.2f é diferente do número 2: %.2f \n", numero1, numero2)
	}
}

func compararTamanhoNomes() {
	var nome1 string
	var nome2 string

	fmt.Print("Digite um nome: ")
	fmt.Scan(&nome1)

	fmt.Print("Digite outro nome: ")
	fmt.Scan(&nome2)

	if len(nome1) > len(nome2) {
		fmt.Printf("O número de caracters do nome 1 é maior que o do nome 2 \n")
	}

	if len(nome1) < len(nome2) {
		fmt.Printf("O número de caracters do nome 1 é menor que o do nome 2 \n")
	}

	if len(nome1) == len(nome2) {
		fmt.Printf("O tamanho de caracters do nome 1 é igual que o do nome 2 \n")
	}

	if len(nome1) != len(nome2) {
		fmt.Printf("O tamanho de caracters do nome 1 é maior que o do nome 2 \n")
	}
}

func compararNomes() {
	var nome1 string
	var nome2 string

	fmt.Print("Digite um nome: ")
	fmt.Scan(&nome1)

	fmt.Print("Digite outro nome: ")
	fmt.Scan(&nome2)

	if nome1 == nome2 {
		fmt.Printf("O nome 1: %s é igual ao nome 2: %s \n", nome1, nome2)
	}

	if nome1 != nome2 {
		fmt.Printf("O nome 1: %s é diferente do nome 2: %s \n", nome1, nome2)
	}
}
