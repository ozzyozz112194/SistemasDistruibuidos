package main

import (
	"fmt"
)

func triangleArea(height int, base int) int {
	var divBase int = base / 2
	return divBase * height
}

func main() {

	fmt.Print("Ingresa la altura \n")
	var height int
	fmt.Scanln(&height)
	fmt.Print("Ingresa la base\n")
	var base int
	fmt.Scanln(&base)

	println("Resultado es  ", triangleArea(height, base))
}
