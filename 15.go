package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

/*
Проблема заключается в нерациональном использованиие памяти:
храним строку длиной 1024 из которых используется всего 100
*/
var justString string

// функция создает строку из рандомных букв
func createHugeString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var result string

	for i := 0; i < length; i++ {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			panic(err)
		}
		result += string(charset[randomIndex.Int64()])
	}
	return result
}

func someFunc() {
	v := createHugeString(1 << 10) //создаем строку
	buffer := make([]byte, 100)    //создаем буфер из 100 элементов
	copy(buffer, v[:100])          //копируем первые 100 элементов исходной строки в буфер
	justString = string(buffer)    //теперь justString не ссылается на исходную строку
}

func main() {
	someFunc()
	fmt.Println(justString)

}
