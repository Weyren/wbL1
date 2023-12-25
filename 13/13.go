package main

import "fmt"

func main() {
	a := 101
	b := 10
	fmt.Println(a, b)

	a = a - b    //разница между a и b
	b = b + a    //b складываем с разницей между a и b: b + a - b = a
	a = -(a - b) //из разницы между a и b вычитаем b(=a) и меняем знак: -(a - b - a) = -(-b) = b
	fmt.Println(a, b)

}
