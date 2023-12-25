package main

//import (
//	"context"
//	"fmt"
//	"sync"
//	"time"
//)
//
//func reader(ctx context.Context, dataChan <-chan int, wg *sync.WaitGroup) {
//	defer wg.Done()
//	for {
//		select {
//		case data, ok := <-dataChan:
//			if !ok {
//				return
//			}
//			fmt.Println(data)
//		case <-ctx.Done():
//			return
//		}
//	}
//}
//
//func writer(ctx context.Context, dataChan chan<- int, wg *sync.WaitGroup) {
//	defer close(dataChan)
//	defer wg.Done()
//	i := 0
//	for {
//		select {
//		case dataChan <- i:
//			i++
//			time.Sleep(time.Millisecond)
//		case <-ctx.Done():
//			return
//		}
//	}
//}
//func main() {
//	var wg sync.WaitGroup
//
//	startTime := time.Now()
//	timeOut := time.Duration(2 * time.Second)
//	dataChan := make(chan int)
//	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
//	defer cancel()
//	wg.Add(2)
//	go reader(ctx, dataChan, &wg)
//	go writer(ctx, dataChan, &wg)
//	<-ctx.Done()
//	wg.Wait()
//	endTime := time.Since(startTime)
//	fmt.Println("time:", endTime)
//}
