package main

//import "fmt"
//
//func cross(a, b map[int]*struct{}) map[int]*struct{} {
//	c := make(map[int]*struct{})
//	for keyA, _ := range a {
//		for keyB, _ := range b {
//			if keyA == keyB {
//				c[keyA] = &struct{}{}
//			}
//		}
//	}
//	return c
//}
//
//func main() {
//	a := make(map[int]*struct{})
//
//	b := make(map[int]*struct{})
//
//	a[9] = &struct{}{}
//	a[5] = &struct{}{}
//	a[0] = &struct{}{}
//	a[0] = &struct{}{}
//
//	b[1] = &struct{}{}
//	b[2] = &struct{}{}
//	b[0] = &struct{}{}
//
//	c := cross(a, b)
//	fmt.Println(a)
//	fmt.Println(c)
//}
