package main

import (
	"fmt"
	"os"
)

func squareArea(width int, height int) int {
	return width + height
}

func circleArea(radius float64) float64 {
	const pi float64 = 3.14
	var radiusSquare float64 = radius * radius

	return pi * radiusSquare
}

func triangleArea(height int, base int) int {
	var divBase int = base / 2
	return divBase * height
}

func conversion(fahrenheit int) int {
	return (fahrenheit - 32) * 5 / 9

}

func menu() {

	var opcion int

	fmt.Scanf("%d", &opcion)

	switch opcion {
	case 1:
		fmt.Print("Ingresa la temp \n")
		var temp int
		fmt.Scanln(&temp)
		println("Your result is : ", conversion(temp))

	case 2:
		var width int
		var height int
		println("Your result is : ", squareArea(width, height))

	case 3:
		var rad float64
		println("Your result is : ", circleArea(rad))

	case 4:
		var height int
		var base int
		println("Your result is : ", triangleArea(height, base))

	case 5:
		println("exit")
		os.Exit(5)

	case 6:
		println("Invalid input")

	}

}

func main() {
	println("\n1)Conversion\n2)Area of square\n3)Area of circle\n4)Area of triangle\n5)Exit\nOption : ")

	for {
		menu()
	}
}
