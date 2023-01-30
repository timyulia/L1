package main

import "fmt"

func reverse(str string) string {
	r := []rune(str) // слайс рун
	n := len(r)
	for i := 0; i < n/2; i++ {
		r[i], r[n-i-1] = r[n-i-1], r[i] // симметричные элементы обмениваются местами
	}
	return string(r)
}

func main() {
	str := "adfwgf"
	fmt.Println(reverse(str))
}
