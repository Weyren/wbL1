package main

import "fmt"

func quickSort(arr []int) {
	if len(arr) < 2 { //когда массив меньше 2 элементов выходим из функции
		return
	}

	pivot := arr[len(arr)/2] //опорный элемент - середина массива
	left := make([]int, 0)   //для значений < опорного элемента
	right := make([]int, 0)  //для значений > опорного элемента
	equal := make([]int, 0)  //для значений = опорному элементу

	for _, value := range arr { //итерируемся по исходному массиву
		switch { //добавляем меньшие равные и большие опорного элемента в соответствующие слайсы
		case value < pivot:
			left = append(left, value)
		case value == pivot:
			equal = append(equal, value)
		case value > pivot:
			right = append(right, value)
		}
	}

	quickSort(left)  //вызываем сортировку для левой и правой части
	quickSort(right) //и проваливаемся вглубь рекурсии
	//собираем слайс
	copy(arr, append(append(left, equal...), right...)) // ... - шорткат для передачи слайса по элементам
}

func main() {
	arr := []int{3, 14, 1, 7, 9, 8, 11, 6, 4, 2}
	quickSort(arr)
	fmt.Println(arr)
}
