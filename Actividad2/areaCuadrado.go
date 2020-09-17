package main

import (
	"fmt"
)

func squareArea(width int, height int) int {
	return width + height
}

func main() {
	fmt.Print("Ingresa el ancho\n")
	var width int
	fmt.Scanln(&width)

	fmt.Print("Ingresa el altura \n")
	var height int
	fmt.Scanln(&height)

	println("Resultado es ", squareArea(width, height))
}
