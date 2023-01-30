package main

import "fmt"

func inter(set1, set2 map[int]struct{}) map[int]struct{} {
	set := make(map[int]struct{})
	for key := range set1 { // иттерации по первому множеству
		if _, ok := set2[key]; ok { //проверка наличия такого элемента во втором множестве
			set[key] = struct{}{} //добавление в результат
		}
	}
	return set
}

func main() {
	set1 := map[int]struct{}{1: {}, 2: {}, 4: {}, 7: {}, 89: {}, 15: {}}
	set2 := map[int]struct{}{1: {}, 5: {}, 4: {}, 7: {}, 89: {}}
	fmt.Println(inter(set1, set2))
}
