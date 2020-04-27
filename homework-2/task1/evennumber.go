//1. Написать функцию, которая определяет, четное ли число.
package main

import "fmt"

func evenNumber(num int) {
	if num%2 == 0 {
		fmt.Println(num, "четное число")
	} else {
		fmt.Println(num, "нечетное число")
	}
}

func main() {
	fmt.Println("Привет. Это программа для определения, четное ли число.\nВведите число")
	var anyNumber int
	fmt.Scanln(&anyNumber)
	evenNumber(anyNumber)
}
