package main

//import "fmt"
//
//func main() {
//	words := []string{"cat", "cat", "dog", "cat", "tree"} //исходная строка
//	set := make(map[string]*struct{})                     //мапа с ключом - стркой и значением - ссылкой на пустую структуру
//	//это удобно ловить ошибки при обращении по несуществуюшему ключу
//	//пример: set["cat"] == &{} - значение есть
//	//set["wolf"] == <nil> - значения нет
//	for _, word := range words {
//		set[word] = &struct{}{}
//	}
//
//	for key, _ := range set { //вывод множества
//		fmt.Print(key, " ")
//	}
//
//}
