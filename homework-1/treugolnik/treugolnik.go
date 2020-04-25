package main

//Даны катеты прямоугольного треугольника. Найти его площадь, периметр и гипотенузу. Используйте тип данных float64 и функции из пакета math.

import (
	"fmt"
	"math"
)

func main() {
	var sideA float64 = 10
	var sideB float64 = 20
	var hypotenuseTriangle float64
	var perimeterTriangle float64
	var areaTriangle float64
	areaTriangle = sideA * sideB / 2
	fmt.Println("Даны катеты прямоугольного треугольника, сторона A =", sideA, "сторона B =", sideB, "\nС такими параметрами:")
	fmt.Printf("Площадь треугольника = %.2f \n", areaTriangle)
	hypotenuseTriangle = math.Sqrt(math.Pow(sideA, 2) + math.Pow(sideB, 2))
	fmt.Printf("Гипотенуза треугольника или по другому сторона C = %.2f \n", hypotenuseTriangle)
	perimeterTriangle = sideA + sideB + hypotenuseTriangle
	fmt.Printf("Периметр треугольника = %.2f \n", perimeterTriangle)
}
