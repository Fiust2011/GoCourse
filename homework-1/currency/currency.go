package main

import "fmt"

func main() {
	var courseDoll = 74
	var sumRub int
	var sumDoll int
	fmt.Println("Привет. Это программа для конвертации рублей в доллары\nКакую сумму в рублях нужно конвертировать в доллары?")
	fmt.Scanln(&sumRub)
	sumDoll = sumRub / courseDoll
	fmt.Printf("Результат: %x \n", sumDoll)
}
