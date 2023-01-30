package main

import (
	"fmt"
	"sync"
)

func mult(ch1, ch2 chan int) {
	for val := range ch1 { //читает из первого канала
		ch2 <- val * 2 //записывает произведение во второй канал
	}
	close(ch2) //закрывает второй канал
}

func printer(ch2 chan int, wg *sync.WaitGroup) {
	for val := range ch2 { //читает из второго канала
		fmt.Println(val) //печатает
	}
	wg.Done() //обнуляет счетчик
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	ch1 := make(chan int, len(nums)) //канал для чисел из массива
	ch2 := make(chan int, len(nums)) //канал для результатов
	go mult(ch1, ch2)                //вычисление произведения
	go printer(ch2, &wg)             //печать
	for i := 0; i < len(nums); i++ {
		ch1 <- nums[i] //закрывает первый канал
	}
	close(ch1)
	wg.Wait() //ожижание завершения печати

}
