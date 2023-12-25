package main

import "time"

func customSleep(d time.Duration) {
	<-time.After(d) //через промежуток d в канал произойдет чтение функция завершится горутина продолжит работу
}

func main() {
	customSleep(5 * time.Second)
}
