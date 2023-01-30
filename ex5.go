package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	var duration time.Duration
	duration = 5 //5 секунд по умолчанию
	if len(os.Args) > 1 {
		d, _ := strconv.Atoi(os.Args[1]) //чтение из командной строки количество секунд
		duration = time.Duration(d)
	}
	values := make(chan int) //канал для данных
	defer close(values)
	t := time.NewTimer(duration * time.Second) //запуск таймера на N секунд
	go func() {
		for {
			fmt.Println(<-values) //печать значений
		}
	}()
	for {
		select {
		case <-t.C: //остановка таймера
			return
		case values <- rand.Intn(1000): //запись данных в канал
		}
	}
}
