//2. Написать функцию, которая определяет, делится ли число без остатка на 3.
package main

import "fmt"

func remainderNumber(num int) {
	if num%3 == 0 {
		fmt.Printf("Число %v делится нацело на 3 \n", num)
	} else {
		fmt.Printf("Число %v не делится нацело на 3 \n", num)
	}
}

func main() {

	fmt.Println("Привет. Это программа, которая определяет, делится ли число без остатка на 3.\nВведите число")
	var anyNumber int
	fmt.Scanln(&anyNumber)
	remainderNumber(anyNumber)
}
