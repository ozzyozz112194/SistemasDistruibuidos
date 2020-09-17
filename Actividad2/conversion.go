package main

import (
	"fmt"
)

func conversion(fahrenheit int) int {
	return (fahrenheit - 32) * 5 / 9

}

func main() {
	var temp int
	println("Ingresa la temperatura")
	fmt.Scanf("%d", &temp)

	println("Resultado es  : C ", conversion(temp))
}
