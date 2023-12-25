package main

import (
	"fmt"
	"strings"
)

func check(str string) bool {
	str = strings.ToLower(str)       //перевод букв в нижний регистр
	set := make(map[int32]*struct{}) //мапа int32(алиас рун) для хранения уникальных символов
	for _, symbol := range str {     //поскольку инерируемся по строке символы конвернируются в int32
		if set[symbol] != nil { //если символ уже встречался значит строка не уникальная
			return false
		} else {
			set[symbol] = &struct{}{} //если символа не было кладем ссылку на структуру
		}
	}
	return true //уникальная
}

func main() {
	var str string
	fmt.Scan(&str)
	isUnique := check(str)
	fmt.Println(isUnique)
}
