package main

import "fmt"

type Viagem struct {
	Origem  string
	Destino string
	Data    string
	Preco   float64
}

func ViagemMaisCara(viagens []Viagem) Viagem {
	var viagemMaisCara Viagem

	for _, v := range viagens {
		if v.Preco > viagemMaisCara.Preco {
			viagemMaisCara = v
		}
	}

	return viagemMaisCara
}

func main() {
	viagens := []Viagem{
		{Origem: "Cidade A", Destino: "Cidade B", Data: "2023-01-01", Preco: 500.0},
		{Origem: "Cidade C", Destino: "Cidade D", Data: "2023-02-01", Preco: 700.0},
		{Origem: "Cidade E", Destino: "Cidade F", Data: "2023-03-01", Preco: 600.0},
	}

	viagemMaisCara := ViagemMaisCara(viagens)
	fmt.Printf("A viagem mais cara é de %s para %s, na data %s, com preço de %.2f\n", viagemMaisCara.Origem, viagemMaisCara.Destino, viagemMaisCara.Data, viagemMaisCara.Preco)
}
