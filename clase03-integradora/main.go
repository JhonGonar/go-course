package main

import (
	"errors"
	"fmt"
)

const (
	dog = "dog"
	cat = "cat"
)

func gramajePerro(cantidad float64) float64 {
	return cantidad * 10000
}
func gramajeGato(cantidad float64) float64 {
	return cantidad * 5000
}

func animal(especie string) (func(cantidadAnimales float64) float64, error) {
	switch especie {
	case dog:
		return gramajePerro, nil
	case cat:
		return gramajeGato, nil
	default:
		return nil, errors.New("No es un animal valido")
	}

}
func main() {
	animalDog, err1 := animal(dog)
	animalCat, err2 := animal(cat)
	variable0, err3 := animal("hamsters")

	fmt.Println(err1)
	fmt.Println(err2)
	fmt.Println(err3)

	var amount float64
	amount += animalDog(10)
	amount += animalCat(10)
	amount += variable0(15)
	fmt.Println(amount, "gramos de comida son necesarios.")
}
