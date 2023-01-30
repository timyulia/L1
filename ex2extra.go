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
	for i := 0; i < len(nums); i++ {
		fmt.Println(<-ch) //чтение из канала и вывод
	}
}
