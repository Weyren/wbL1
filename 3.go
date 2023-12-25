package main

import "fmt"

// функция пишет в канал квадрат полученного числа
func square(i int, nums chan int) {
	nums <- i * i
}

func main() {
	nums := make(chan int, 5) //создали канал целочисленный
	array := [5]int{2, 4, 6, 8, 10}
	sum := 0

	for _, value := range array {
		go square(value, nums)
	}
	for i := 0; i < 5; i++ {
		select {
		case num := <-nums:
			sum += num
		}
	}
	fmt.Println(sum, "\n")
}
