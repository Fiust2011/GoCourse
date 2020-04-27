package main

import "fmt"

//3.Написать функцию, которая последовательно выводит на экран 100 первых чисел Фибоначчи, начиная от 0.
//Числа Фибоначчи определяются соотношениями Fn=Fn-1 + Fn-2.
func fibonacci() func() int {
	numbers := make(map[int]int)
	n := 0
	return func() int {
		if n == 0 {
			numbers[n] = 0
			n++
			return 0
		}
		if n == 1 {
			numbers[n] = 1
			n++
			return 1
		}
		number := numbers[n-1] + numbers[n-2]
		numbers[n] = number
		n++
		return number
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 100; i++ {
		fmt.Println(f())
	}
}
