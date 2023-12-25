package main

////map+RWMutex (вариант ускоряет работу в случае если необходимо конкурентное чтение во время записи в мапу
////RWMutex есть Rlock() и RUnlock() => возможно конкурентное чтение при заблокированном мьютексе
//import (
//	"fmt"
//	"sync"
//)
//
//// структура хранилища
//type Storage struct {
//	sync.RWMutex
//	m map[int]int
//}
//
//// конструктор
//func NewStorage() *Storage {
//	return &Storage{
//		m: make(map[int]int),
//	}
//}
//
//// метод добавления значения
//func (s *Storage) Store(key, value int) {
//	s.Lock()         //лочим мьютекс(то же самое что и s.Mutex.Lock() но применил встраивание
//	defer s.Unlock() // разлочить мьютекс после действий с хранилищем
//	s.m[key] = value //запись значиния
//}
//
//func main() {
//	var wg sync.WaitGroup //wg для синхронизаци горутин
//
//	s := NewStorage() //создаем объект хранилища
//	for i := 0; i < 10; i++ {
//		i := i //чтобы не потерять значение
//		wg.Add(1)
//		go func(i int) {
//			s.Store(i, i)
//			wg.Done()
//		}(i)
//	}
//	wg.Wait()
//	fmt.Println(s.m)
//}
