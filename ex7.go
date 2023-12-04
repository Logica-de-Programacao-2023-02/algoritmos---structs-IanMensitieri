package main

import "fmt"

type Animal struct {
	Nome   string
	Especie string
	Idade  int
	Som    string
}

func ModificarSom(a *Animal, novoSom string) {
	a.Som = novoSom
}

func ImprimirInfoAnimal(a Animal) {
	fmt.Printf("Nome: %s\n", a.Nome)
	fmt.Printf("Espécie: %s\n", a.Especie)
	fmt.Printf("Idade: %d anos\n", a.Idade)
	fmt.Printf("Som que faz: %s\n", a.Som)
}

func main() {
	animalExemplo := Animal{
		Nome:   "Fido",
		Especie: "Cachorro",
		Idade:  5,
		Som:    "Latido",
	}

	ImprimirInfoAnimal(animalExemplo)

	ModificarSom(&animalExemplo, "Miado")

	ImprimirInfoAnimal(animalExemplo)
}
