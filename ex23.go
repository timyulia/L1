package main

import "fmt"

func remove(i int, nums []int) []int {
	return append(nums[:i], nums[i+1:]...) // к слайсу с элеменами до  i-го добавить элементы после i-го
}

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	nums = remove(9, nums)
	fmt.Println(nums)
}
