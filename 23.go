package main

import "fmt"

func del(sl *[]int, pos int) {
	*sl = append((*sl)[:pos], (*sl)[pos+1:]...)
}

func main() {
	sl := []int{1, 2, 3, 4, 5}
	fmt.Println(sl)
	del(&sl, 4)
	fmt.Println(sl)
}
