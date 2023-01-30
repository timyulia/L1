package main

import (
	"fmt"
	"math"
)

func equalOne(n int64, i int) bool {
	del := int64(math.Pow(float64(2), float64(i))) //какому числу соответсвует данный бит
	return n%(del*2)/del == 1
}

func equalOne2(n int64, i int) bool {
	return (n>>(i-1))%2 == 1 //сдвиг
}

func change(n int64, i, val int) int64 { //уставновить i-ый бит числа n в значение val
	if equalOne(n, i) && val == 0 { //если 1, а нужен 0
		n = n - int64(math.Pow(float64(2), float64(i))) //вычесть 2 в степени i
		return n
	}
	if !equalOne(n, i) && val == 1 { //если 0, а нужен 1
		n = n + int64(math.Pow(float64(2), float64(i))) //прибавить 2 в степени i
	}
	return n
}

func main() {
	// биты нумеруются справа налево с нуля
	fmt.Println(change(3882, 10, 0)) //2858=101100101010 3882=111100101010
	fmt.Println(change(3882, 10, 1)) //3882
	fmt.Println(change(223, 3, 0))   //215=11010111 223=11011111
	fmt.Println(change(5, 1, 1))     //7=111
}
