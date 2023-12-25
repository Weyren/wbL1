package main

//import (
//	"fmt"
//	"math/rand"
//)
//
//func setBit(num *int64, position int64, value int) error {
//
//	if position < 0 || position > 64 {
//		return fmt.Errorf("position can't be less than 0 and bigger than 64")
//	}
//
//	var mask int64 = 1 << position //маска для установки в необходимую позицию(нумерация с 0)
//	if value == 1 {
//		fmt.Printf("%064b\n ", mask)
//		*num = *num | mask //поразрядная дизъюнкция (ИЛИ) - бит равен 1 если соответствующий бит маски равен 1
//	} else if value == 0 {
//		fmt.Printf("%064b\n", mask)
//		*num = *num &^ mask //сброс бита (И НЕ) - бит равен 0 если соответствующий бит маски равен 1
//	} else {
//		return fmt.Errorf("value can be only 0 or 1")
//	}
//
//	return nil
//}
//
//func main() {
//
//	var num int64
//	num = -rand.Int63()                //генерация рандомного числа
//	fmt.Printf("%064b %d\n", num, num) //вывод до
//
//	err := setBit(&num, 0, 1) //передаем указатель на число чтобы в функции менять его
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Printf("%064b %d", num, num) //вывод после
//
//}
