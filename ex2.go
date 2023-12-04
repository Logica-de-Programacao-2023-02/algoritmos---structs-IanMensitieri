package main

import "fmt"

type Endereco struct {
	Rua     string
	Numero  int
	Cidade  string
	Estado  string
}

type Pessoa struct {
	Nome    string
	Idade   int
	Endereco Endereco
}

func ImprimirEnderecoCompleto(p Pessoa) {
	fmt.Printf("Endereço de %s:\n", p.Nome)
	fmt.Printf("Rua: %s, Número: %d, Cidade: %s, Estado: %s\n", p.Endereco.Rua, p.Endereco.Numero, p.Endereco.Cidade, p.Endereco.Estado)
}

func main() {
	enderecoExemplo := Endereco{Rua: "Rua ABC", Numero: 123, Cidade: "CidadeXYZ", Estado: "EstadoABC"}
	pessoaExemplo := Pessoa{Nome: "João", Idade: 30, Endereco: enderecoExemplo}

	ImprimirEnderecoCompleto(pessoaExemplo)
}
