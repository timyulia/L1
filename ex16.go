package main

import (
	"fmt"
)

func sort(nums []int) []int {

	if len(nums) < 2 {
		return nums
	}
	mid := nums[len(nums)/2] //выбираем опорный элемент
	var left, right, equal []int
	for i := 0; i < len(nums); i++ {
		if mid < nums[i] {
			right = append(right, nums[i]) //часть с большими числами
		} else if mid > nums[i] {
			left = append(left, nums[i]) //часть с меньшим числами
		} else {
			equal = append(equal, nums[i]) //равные
		}
	}

	res := append(sort(left), equal...) //к меньшим слева дописываем равные справа
	res = append(res, sort(right)...)   //затем большие

	return res
}

func main() {
	nums := []int{-3, 6, -10, 23, 600, 3, -90}
	fmt.Println(sort(nums))
}
