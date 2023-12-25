package main

import (
	"fmt"
	"sync"
)

// функция вывода квадрата числа в stdout
func square(i int, wg *sync.WaitGroup) {
	fmt.Print(i*i, " ")
	wg.Done() //уменьшаем счетчик
}

func main() {

	var wg sync.WaitGroup //объявляем WaitGroup для синхронизации горутин

	array := [5]int{2, 4, 6, 8, 10}

	for _, value := range array {
		wg.Add(1)             // увеличиваем счетчик wg
		go square(value, &wg) // передаем в функцию число и ссылку на wg
		// передача по значению не подходит т.к. она приведет к deadlock(wg не дождется)
	}
	wg.Wait() //ждем выполнения всех горутин
	fmt.Print("\n")

}
