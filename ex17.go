package main

import "fmt"

func search(nums []int, target int) int {
	a := 0
	b := len(nums) - 1
	var mid int
	for a <= b {
		mid = (a + b) / 2 //вычисляем средний индекс
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			a = mid + 1 //сдвигаемся вправо
		} else {
			b = mid - 1 //сдвигаемся влево
		}
	}
	return -1
}

func main() {
	var nums []int = []int{-6, -4, 7, 8, 9, 12, 18}
	fmt.Println(search(nums, -8)) //-1 если нет, иначе место
}
