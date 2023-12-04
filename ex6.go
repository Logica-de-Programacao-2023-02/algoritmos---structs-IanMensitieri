package main

import (
	"fmt"
	"time"
)

type Funcionario struct {
	Nome    string
	Salario float64
	Idade   int
}

func AumentarSalario(f *Funcionario, percentual float64) {
	aumento := f.Salario * percentual / 100
	f.Salario += aumento
}

func DiminuirSalario(f *Funcionario, percentual float64) {
	desconto := f.Salario * percentual / 100
	f.Salario -= desconto
}

func TempoDeServico(f Funcionario) int {
	idadeInicioTrabalho := 18
	anosDeServico := time.Now().Year() - (18 + f.Idade)
	return anosDeServico
}

func main() {
	funcionarioExemplo := Funcionario{
		Nome:    "João",
		Salario: 5000.0,
		Idade:   30,
	}

	fmt.Printf("Salário inicial de %s: %.2f\n", funcionarioExemplo.Nome, funcionarioExemplo.Salario)

	AumentarSalario(&funcionarioExemplo, 10.0)
	fmt.Printf("Novo salário após aumento: %.2f\n", funcionarioExemplo.Salario)

	DiminuirSalario(&funcionarioExemplo, 5.0)
	fmt.Printf("Novo salário após desconto: %.2f\n", funcionarioExemplo.Salario)

	tempoServico := TempoDeServico(funcionarioExemplo)
	fmt.Printf("Tempo de serviço de %s: %d anos\n", funcionarioExemplo.Nome, tempoServico)
}
