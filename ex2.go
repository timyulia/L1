package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup //чтобы main не завершилась раньше остальных горутин
	for i := 0; i < len(nums); i++ {
		wg.Add(1)
		go func(n int) {
			fmt.Printf("%d - %d\n", n, n*n)
			wg.Done()
		}(nums[i]) //передача функции текущего элемента массива
	}
	wg.Wait()
}
