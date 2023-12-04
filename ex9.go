package main

import (
	"fmt"
)

type Aluno struct {
	Nome  string
	Idade int
	Notas []float64
}

func AdicionarNota(a *Aluno, nota float64) {
	a.Notas = append(a.Notas, nota)
}

func RemoverNota(a *Aluno, indice int) {
	if indice >= 0 && indice < len(a.Notas) {
		a.Notas = append(a.Notas[:indice], a.Notas[indice+1:]...)
	}
}

func CalcularMedia(a Aluno) float64 {
	if len(a.Notas) == 0 {
		return 0.0
	}

	total := 0.0
	for _, nota := range a.Notas {
		total += nota
	}

	return total / float64(len(a.Notas))
}

func ImprimirInfoAluno(a Aluno) {
	fmt.Printf("Nome: %s\n", a.Nome)
	fmt.Printf("Idade: %d anos\n", a.Idade)
	fmt.Printf("Média das notas: %.2f\n", CalcularMedia(a))
}

func main() {
	alunoExemplo := Aluno{
		Nome:  "João",
		Idade: 20,
		Notas: []float64{7.5, 8.0, 9.5},
	}

	ImprimirInfoAluno(alunoExemplo)

	AdicionarNota(&alunoExemplo, 8.5)
	ImprimirInfoAluno(alunoExemplo)

	RemoverNota(&alunoExemplo, 2)
	ImprimirInfoAluno(alunoExemplo)
}
