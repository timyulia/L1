package main

import "fmt"

func main() {
	ch := make(chan int) //канал для передачи результата
	nums := []int{2, 4, 6, 8, 10}
	for _, val := range nums {
		go func(val int, ch chan int) {
			ch <- val * val //запись квадрата в канал
		}(val, ch)
	}
	res := 0
	for i := 0; i < len(nums); i++ {
		res += <-ch //чтение из канала и сложение
	}
	fmt.Println(res)
}
