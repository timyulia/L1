package main

import (
	"fmt"
	"strings"
)

func someFunc(justString *string) {
	v := createHugeString(1 << 10)
	*justString = strings.Clone(v[:100]) //копируем в новое место памяти, чтобы сборщик мусора смог почистить большую строку
}

func createHugeString(n int) string { // функция создания большрй строки
	var b strings.Builder
	filler := "HugeString"
	extra := n % len(filler) //сколько лишних символов
	for i := extra; i < n; i += 10 {
		b.WriteString(filler)
	}
	for i := 0; i < extra; i++ { //заполнение лишних символов
		b.WriteByte('_')
	}
	return b.String()
}

func main() {
	var justString string // создаем локальную переменную
	someFunc(&justString) //передаем указатель на строку, чтобы получить изменение
	fmt.Println(justString)
}
