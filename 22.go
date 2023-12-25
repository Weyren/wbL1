package main

import (
	"fmt"
	"math/big"
)

func calc(a, b *big.Int, operator string) *big.Int {
	res := new(big.Int) //новый указатель на big.Int для записи результата
	switch operator {
	case "+":
		return res.Add(a, b)
	case "-":
		return res.Sub(a, b)
	case "*":
		return res.Mul(a, b)
	case "/":
		return res.Div(a, b)
	default:
		return res
	}

}

func main() {
	a := new(big.Int) //создаем два указателя на big.Int
	b := new(big.Int)

	a.SetInt64(10000000000000000) //присваиваем значения
	b.SetInt64(222222222222222)

	fmt.Println(calc(a, b, "+"))
	fmt.Println(calc(a, b, "-"))
	fmt.Println(calc(a, b, "*"))
	fmt.Println(calc(a, b, "/"))

}
