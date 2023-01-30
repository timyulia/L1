package main

import (
	"fmt"
	"strings"
)

func rotate(str string) string {
	words := strings.Split(str, " ") //делим строку на отделные слова
	n := len(words)
	for i := 0; i < n/2; i++ {
		words[i], words[n-1-i] = words[n-1-i], words[i] //меняем попарно (первый - последний, второй - предпоследний)
	}
	return strings.Join(words, " ") // создаем строку с пробелами
}

func main() {
	str := "snow dog sun blue light sunday"
	fmt.Println(rotate(str))

}
