package main

//import "fmt"
//
//func reverse(str string) string {
//	res := []rune(str) //слайс рун по исходной строке
//	lenght := len(res) - 1
//	for i := 0; i < lenght/2; i++ {
//		res[i], res[lenght-i] = res[lenght-i], res[i] //меняем местами крайние элементы в доступной области
//	}
//	return string(res) //приведение слайса рун к строке
//}
//
//func main() {
//	var str string
//	_, err := fmt.Scan(&str) //ввод с консоли
//	if err != nil {
//		fmt.Println("Enter string")
//	}
//
//	str = reverse(str)
//	fmt.Println(str)
//}
