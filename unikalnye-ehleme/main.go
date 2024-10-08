package main

import "fmt"

func main() {
	var sizeArr int
	fmt.Scan(&sizeArr)

	// Создаем мапу
	arrMap := mapCreate(sizeArr)

	// Подсчитываем уникальные числа и выводим результат
	fmt.Println(countUnique(arrMap))
}

// Функция для создания мапы
func mapCreate(sizeArr int) map[int]int {
	var num int
	arrMap := make(map[int]int)

	for i := 0; i < sizeArr; i++ {
		fmt.Scan(&num)
		// Считаем количество вхождений каждого числа
		arrMap[num]++
	}
	return arrMap
}

// Функция для подсчета уникальных чисел
func countUnique(arrMap map[int]int) int {
	tempUnique := 0
	for _, count := range arrMap {
		if count == 1 { // встречается ли число только 1 раз
			tempUnique++
		}
	}
	return tempUnique
}
