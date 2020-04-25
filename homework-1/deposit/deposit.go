package main

//* Пользователь вводит сумму вклада в банк и годовой процент. Найти сумму вклада через 5 лет.

import (
	"fmt"
	"math"
)

func main() {
	var futureValue float64
	var initialAmount float64
	var interestRate float64
	var numberOfYears float64 = 5
	fmt.Println("Привет. Это программа для вычисления суммы вклада, через 5 лет.\nВведите первоначальну сумму вклада")
	fmt.Scanln(&initialAmount)
	fmt.Println("Введите годовую процентную ставку")
	fmt.Scanln(&interestRate)
	futureValue = initialAmount * math.Pow(1+interestRate/100, numberOfYears)
	fmt.Printf("Сумма вклада через 5 лет = %.3f\n", futureValue)
}
