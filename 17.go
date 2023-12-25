package main

import (
	"fmt"
	"sort"
)

func search(arr []int, value int) (int, error) {
	left := 0             //начало области поиска
	right := len(arr) - 1 //конец области поиска
	for left < right {
		mid := (left + right) / 2 //среднее значение
		switch {
		case arr[mid] < value:
			left = mid + 1 // ищем от mid+1 до left
		case arr[mid] > value:
			right = mid - 1 //ищем от left до  mid - 1
		case arr[mid] == value:
			return mid, nil //ответ
		}
	}
	return -1, fmt.Errorf("value not found")
}

func main() {
	arr := []int{1, 132, 443, 492, 223, 3, 4, 8}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j] //функция сортировки по возрастанию
	})

	a, err := search(arr, 443)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(arr)
	fmt.Println("position:", a)
}
