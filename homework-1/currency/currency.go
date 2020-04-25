package main

import "fmt"

func main() {
	var courseDoll float64 = 74
	var sumRub float64
	var sumDoll float64
	fmt.Println("Привет. Это программа для конвертации рублей в доллары\nКакую сумму в рублях нужно конвертировать в доллары?")
	fmt.Scanln(&sumRub)
	sumDoll = sumRub / courseDoll
	fmt.Printf("Результат: %x \n", sumDoll)
}
