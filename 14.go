package main

import (
	"fmt"
	"reflect"
)

func vartype(v interface{}) {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Int:
		fmt.Println("int", v)
	case reflect.String:
		fmt.Println("string", v)
	case reflect.Bool:
		fmt.Println("bool", v)
	case reflect.Chan:
		fmt.Println("chan", v)
	default:
		fmt.Println("unknown", v)
	}
}

func main() {
	var intVar interface{} = 1
	var stringVar interface{} = "string"
	var chanVar interface{} = make(chan int)
	var boolVar interface{} = true
	var floatVar interface{} = 3.14

	vartype(intVar)
	vartype(stringVar)
	vartype(chanVar)
	vartype(boolVar)
	vartype(floatVar)
}
