package main

import (
	"fmt"
)

func same_params_type(var1, var2 string) {
	println(var1, var2)
}
func main() {
	fmt.Printf("Esto es un número %T", 231)

	var miEntero int16 = 8
	fmt.Printf("Mi entero de 16 bits %d", miEntero)
	same_params_type("Ho", "la")
}
