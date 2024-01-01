package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, id int, dataChannel chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case data, ok := <-dataChannel:
			if !ok {
				fmt.Printf("Worker %d does not received data\n", id)
				return
			}
			fmt.Printf("Worker %d received data: %s\n", id, data)
		case <-ctx.Done():
			fmt.Printf("Worker %d received context cancellation\n", id)
			return
		}
	}
}

func main() {
	// Количество воркеров
	numWorkers := 3

	// Создаем контекст с отменой
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Создаем канал для передачи данных
	dataChannel := make(chan string)

	// Создаем WaitGroup для ожидания завершения всех воркеров
	var wg sync.WaitGroup

	// Запускаем воркеры
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i, dataChannel, &wg)
	}

	// Обработка сигнала Ctrl+C для завершения программы
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signalChannel
		fmt.Println("\nCtrl+C received. Cancelling the context...")
		cancel() // Отменяем контекст для оповещения воркеров о завершении
		wg.Wait()

		return
	}()

	// Главный поток записывает данные в канал
	//for i := 1; i <= 1000; i++ {
	//	dataChannel <- fmt.Sprintf("Data %d", i)
	//	time.Sleep(time.Millisecond * 100)
	//}

	i := 0
L:
	for {
		select {
		case <-ctx.Done():
			break L
		default:
			dataChannel <- fmt.Sprintf("Data %d", i)
			time.Sleep(time.Second)

		}
	}

	fmt.Println("All workers have exited. Program terminated.")
}
