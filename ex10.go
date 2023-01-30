package main

import "fmt"

func det(temp float32) int { //определение группы
	n := int(temp)
	return n / 10 * 10
}

func devide(temp []float32) map[int][]float32 {
	groups := make(map[int][]float32) //хэш-таблица с ключом в виде группы и слайсом с элементами группы
	for _, val := range temp {
		group := det(val)
		groups[group] = append(groups[group], val)
	}
	return groups
}

func main() {
	temps := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := devide(temps)
	for key, value := range groups {
		fmt.Printf("%d: %v  ", key, value)
	}
}
