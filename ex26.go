package main

import (
	"fmt"
	"strings"
)

func check(str string) bool {
	str = strings.ToLower(str)    //все символы к одному регистру (нижнему)
	in := make(map[rune]struct{}) //мэп со значеним в виде пустой структуры, так как оно не нужно
	for _, ch := range str {      //иттерация по строчке
		if _, ok := in[ch]; ok { //если уже есть такой ключ в мэпе, то символ не уникальный
			return false
		}
		in[ch] = struct{}{} //иначе он встретился первый раз, и его надо занести в мэп
	}
	return true
}

func main() {
	fmt.Println(check("abCdefAaf"))
	fmt.Println(check("aabcd"))
	fmt.Println(check("string"))
}
