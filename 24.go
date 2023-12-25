package main

import (
	"fmt"
	"math"
)

// структура точки с двумя полями
type Point struct {
	x float64
	y float64
}

// конструктор принимает 2 значения и возвращает указатель на структуру
func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

// функция считает расстояние используя функции пакета math
func distance(p1, p2 *Point) float64 {
	d := math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
	return d
}

func main() {
	p1 := NewPoint(1, 5) //создали две структуры Point и получили указатели на них
	p2 := NewPoint(2, 2)
	d := distance(p1, p2)
	fmt.Println(d)
}
