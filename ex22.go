package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)                                                    //новое большое число
	a.SetString("4600000000000000000000000000000000000000000000000", 10) //задаем значение с помощью строки, основание 10
	b := new(big.Int)
	b.SetString("230000000000000000000000000000000000000000000000", 10)
	z := new(big.Int)

	fmt.Println(z.Mul(a, b)) //умножение
	fmt.Println(z.Div(a, b)) //деление
	fmt.Println(z.Add(a, b)) //сложение
	fmt.Println(z.Sub(a, b)) //вычитание
}
