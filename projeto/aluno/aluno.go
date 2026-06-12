package aluno

import "fmt"

type Aluno struct {
	Nome string
	N1   float64
	N2   float64
}

func (a Aluno) Media() float64 {
	return (a.N1 + a.N2) / 2
}

func (a Aluno) Status() string {
	media := a.Media()

	switch {
	case media >= 7:
		return "APROVADO"
	case media >= 5:
		return "RECUPERACAO"
	default:
		return "REPROVADO"
	}
}

func ObterInformacoesAluno() Aluno {
	var aluno Aluno

	fmt.Print("Nome do aluno: ")
	fmt.Scan(&aluno.Nome)

	fmt.Print("Primeira nota do aluno: ")
	fmt.Scan(&aluno.N1)

	fmt.Print("Segunda nota do aluno: ")
	fmt.Scan(&aluno.N2)

	return aluno
}

func (a Aluno) RetornoInformacoesAluno() string {
	return fmt.Sprintf(
		"Nome: %s\nNota 1: %.2f\nNota 2: %.2f\nMedia: %.2f\n",

		a.Nome,
		a.N1,
		a.N2,
		(a.N1+a.N2)/2,
	)
}
