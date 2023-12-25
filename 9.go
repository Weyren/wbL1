package main

import (
	"fmt"
	"time"
)

func gen(nums []int) <-chan int {
	in := make(chan int)

	go func() {
		for _, n := range nums {
			in <- n
		}
		close(in)
	}()
	return in
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * 2
		}
		close(out)
	}()
	return out
}

func printsq(out <-chan int) {
	go func() {
		for n := range out {
			fmt.Println(n)
		}
	}()
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	in := gen(nums)
	out := square(in)
	printsq(out)
	time.Sleep(5 * time.Second)
}
