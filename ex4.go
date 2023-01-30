package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {
	n := 3 //число воркеров по умолчанию
	if len(os.Args) > 1 {
		n, _ = strconv.Atoi(os.Args[1]) //чтение аргумента с числом воркеров
	}
	ch := make(chan int) //канал для передачи чисел
	for i := 0; i < n; i++ {
		go func(i int) { //создание отдельных горутин
			for {
				fmt.Printf("worker: %d  value: %d\n", i, <-ch) //вывод из канала
			}
		}(i) //передача номера вокера
	}
	signals := make(chan os.Signal, 1)                    //канал для обработки сигналов
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM) //прием внешних сигналов
	for {
		select {
		case <-signals: //если ctrl C
			return
		default:
		}
		ch <- rand.Intn(100) //передача случайного числа по каналу
	}
}
