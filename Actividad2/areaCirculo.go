package main

import (
	"fmt"
)

func circleArea(radius float64) float64 {
	const pi float64 = 3.14
	var radiusSquare float64 = radius * radius

	return pi * radiusSquare
}

func main() {
	fmt.Print("Ingresa los radianes \n")
	var rad float64
	fmt.Scanln(&rad)

	println("Resultado es  : ", circleArea(rad))
}
