package main

//import (
//	"bufio"
//	"fmt"
//	"os"
//	"strings"
//)
//
//func main() {
//	reader := bufio.NewReader(os.Stdin) //создали ридер
//	str, _ := reader.ReadString('\n')   //прочитали строку
//	words := strings.Fields(str)        //слайс слов строки
//
//	lenght := len(words) - 1
//	for i := 0; i < lenght/2; i++ {
//		words[i], words[lenght-i] = words[lenght-i], words[i] //меняем местами крайние элементы доступной области
//	}
//	str = strings.Join(words, " ") //собрали строку из слайса с разделителем
//	fmt.Println(str)
//}
