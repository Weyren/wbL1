package main

//import (
//	"fmt"
//	"sync"
//	"sync/atomic"
//)
//
//type Counter struct {
//	count int64
//}
//
//func main() {
//	var wg sync.WaitGroup
//	c := Counter{} //по дефолту count = 0
//	fmt.Println(c.count)
//	for i := 0; i < 1000; i++ {
//		wg.Add(1)
//		go func() {
//			atomic.AddInt64(&c.count, 1) //используем атомик - инкриментим счетчик на атомарном уровне
//			wg.Done()
//		}()
//	}
//	wg.Wait() //дождаться всех горутин
//	fmt.Println(c.count)
//}
