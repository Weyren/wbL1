package main

//import (
//	"context"
//	"fmt"
//	"time"
//)
//
//func gorutine1(ctx context.Context) {
//	for {
//		select {
//		case <-ctx.Done():
//			fmt.Println("gorutine1 finished")
//			return
//		default:
//			time.Sleep(time.Second)
//		}
//	}
//}
//
//func gorutine2(stopCh <-chan struct{}) {
//	for {
//		select {
//		case <-stopCh:
//			fmt.Println("gorutine2 finished")
//			return
//		default:
//			time.Sleep(time.Second)
//		}
//	}
//}
//
//func main() {
//	ctx, cancel := context.WithCancel(context.Background())
//	go gorutine1(ctx)
//	cancel()
//
//	stopChan := make(chan struct{})
//	go gorutine2(stopChan)
//	stopChan <- struct{}{}
//	close(stopChan)
//
//	time.Sleep(time.Second * 10)
//}
