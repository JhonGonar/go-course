package main

import (
	"fmt"
)

func categoriaA(valor float64) float64 {
	return valor * 1.5
}
func categoriaB(valor float64) float64 {
	return valor * 1.2
}
func categoriaC(valor float64) float64 {
	return valor
}
func calcularSalarioTotal(horasTrabajadas float64, categoria string) /*func(valor float64)*/ float64 {
	var subtotalSalario = horasTrabajadas * 40 * 4
	switch categoria {
	case "A":
		return categoriaA(subtotalSalario)
	case "B":
		return categoriaB(subtotalSalario)
	case "C":
		return categoriaC(subtotalSalario)
	default:
		return 0.00
	}
}
func main() {

	var ejemplo = calcularSalarioTotal(14.0, "B")

	fmt.Println(ejemplo)
}
