package main

import "fmt"

func main() {
	addressBook := make(map[string][]int)

	addressBook["Alex"] = []int{89996543210}
	addressBook["Bob"] = []int{89167243812}
	addressBook["Bob"] = append(addressBook["Bob"], 89155243627)

	fmt.Println("Создаем адресну книку:\n", addressBook)

	if number, ok := addressBook["Alex"]; ok {
		fmt.Println(number)
	}

	for name, numbers := range addressBook {
		fmt.Println("Абонент:", name)
		for i, number := range numbers {
			fmt.Printf("\t %v) %v \n", i+1, number)
		}
	}

	// количество элементов
	mapLength := len(addressBook)
	fmt.Println("количество элементов в Адресной книге:\n", mapLength)

	// удаление ключа
	delete(addressBook, "Bob")
	fmt.Println("удаление ключа из Адресной кники:\n", addressBook)

	for name, numbers := range addressBook {
		fmt.Println("Абонент:", name)
		for i, number := range numbers {
			fmt.Printf("\t %v) %v \n", i+1, number)
		}
	}
}
