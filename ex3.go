package main

import "fmt"

type Triangulo struct {
	Base   float64
	Altura float64
}

func CalcularAreaTriangulo(t Triangulo) float64 {
	area := (t.Base * t.Altura) / 2
	return area
}

func main() {
	trianguloExemplo := Triangulo{Base: 5.0, Altura: 8.0}
	area := CalcularAreaTriangulo(trianguloExemplo)
	fmt.Printf("A área do triângulo com base %.2f e altura %.2f é %.2f\n", trianguloExemplo.Base, trianguloExemplo.Altura, area)
}
