package main

import (
	"fmt"
)

type Filme struct {
	Titulo      string
	Diretor     string
	Ano         int
	Avaliacoes  []int
}

func AdicionarAvaliacao(f *Filme, avaliacao int) {
	f.Avaliacoes = append(f.Avaliacoes, avaliacao)
}

func RemoverAvaliacao(f *Filme, indice int) {
	if indice >= 0 && indice < len(f.Avaliacoes) {
		f.Avaliacoes = append(f.Avaliacoes[:indice], f.Avaliacoes[indice+1:]...)
	}
}

func CalcularMediaAvaliacoes(f Filme) float64 {
	if len(f.Avaliacoes) == 0 {
		return 0.0
	}

	total := 0
	for _, avaliacao := range f.Avaliacoes {
		total += avaliacao
	}

	return float64(total) / float64(len(f.Avaliacoes))
}

func ImprimirInfoFilme(f Filme) {
	fmt.Printf("Título: %s\n", f.Titulo)
	fmt.Printf("Diretor: %s\n", f.Diretor)
	fmt.Printf("Ano: %d\n", f.Ano)
	fmt.Printf("Média das avaliações: %.2f\n", CalcularMediaAvaliacoes(f))
}

func main() {
	filmeExemplo := Filme{
		Titulo:     "O Filme",
		Diretor:    "Diretor A",
		Ano:        2022,
		Avaliacoes: []int{4, 5, 3, 4, 5},
	}

	ImprimirInfoFilme(filmeExemplo)

	AdicionarAvaliacao(&filmeExemplo, 4)
	ImprimirInfoFilme(filmeExemplo)

	RemoverAvaliacao(&filmeExemplo, 2)
	ImprimirInfoFilme(filmeExemplo)
}
