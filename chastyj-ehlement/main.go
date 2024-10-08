package main

import "fmt"

func main() {
	var size int
	_, err := fmt.Scan(&size)
	if err != nil {
		fmt.Println("Ошибка ввода")
		return
	}

	fmt.Println(numMax(size))
}

func numMax(size int) int {
	freq := make(map[int]int)
	maxCount := 0
	maxNum := 0

	// Заполняем массив
	for i := 0; i < size; i++ {
		var num int
		_, err := fmt.Scan(&num)
		if err != nil {
			fmt.Println("Ошибка ввода")
			return 0
		}
		freq[num]++

		// Обновляем максимальную частоту и число при необходимости
		if freq[num] > maxCount || (freq[num] == maxCount && num > maxNum) {
			maxCount = freq[num]
			maxNum = num
		}
	}
	return maxNum
}
