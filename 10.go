package main

import (
	"fmt"
	"math"
)

// структура хранилища
type Storage struct {
	m map[int][]float64
}

// конструктор
func NewStorage() *Storage {
	return &Storage{m: make(map[int][]float64)}
}

// добавление по ключу
func (s *Storage) Add(key int, val float64) {
	sl, ok := s.m[key]

	if ok == false {
		s.NewKey(key)
	}

	sl = append(sl, val)
	s.m[key] = sl
}

// создание нового ключа
func (s *Storage) NewKey(key int) {
	s.m[key] = []float64{}
}

func main() {
	s := NewStorage()
	values := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	for _, temp := range values {
		var key int
		if temp >= 0 {
			key = int(math.Floor(temp/10) * 10)
		} else {
			key = int(math.Ceil(temp/10) * 10)
		}
		s.Add(key, temp)
	}

	for key, value := range s.m {
		fmt.Println(key, ":", "{", value, "}")
	}
}
