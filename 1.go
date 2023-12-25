package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

// Встраивание структуры Human в структуру Action
type Action struct {
	Human
}

// Метод структуры Human
func (h *Human) introduce() {
	fmt.Printf("My name is %s. I am %d y.o.", h.Name, h.Age)
}

func main() {
	h := Human{Name: "Pavel", Age: 19}
	a := Action{h}
	a.introduce() //вызов метода introduce у структуры Action
	fmt.Print("\n")
}
