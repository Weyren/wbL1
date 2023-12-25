package main

////sync.Map - применяется в случае когда ключи не обновляются часто и чтений >> записей
//import (
//	"fmt"
//	"sync"
//)
//
//func main() {
//	storage := &sync.Map{} //создали объект sync.Map
//
//	var wg sync.WaitGroup
//
//	for i := 1; i < 10; i++ {
//		i := i
//		wg.Add(1)
//		go func(i int) {
//			storage.Store(i, i) //вызов метода Store - добавляет значение по ключу
//			wg.Done()
//		}(i)
//	}
//	wg.Wait()
//
//	fmt.Println(storage)
//}
