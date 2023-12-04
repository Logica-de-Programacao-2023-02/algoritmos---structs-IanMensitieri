package main

import (
	"fmt"
	"math"
)

type Circulo struct {
	raio float64
}

func CalcularArea(c Circulo) float64 {
	area := math.Pi * c.raio * c.raio
	return area
}

func main() {
	circuloExemplo := Circulo{raio: 3.0}
	area := CalcularArea(circuloExemplo)
	fmt.Printf("A área do círculo com raio %.2f é %.2f\n", circuloExemplo.raio, area)
}
