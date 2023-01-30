package main

import (
	"fmt"
	"time"
)

func sleep1(duration int64) {
	start := time.Now() //запоминаем начальное время
	fmt.Printf("sleeping for %d seconds\n", duration)
	for {
		if start.Unix()+duration <= time.Now().Unix() { //сравниваем начальное плюс продолжительность и текущее время
			fmt.Println("waking up")
			return
		}
	}
}

func sleep2(duration time.Duration) {
	timer := time.NewTimer(duration * time.Second) //запуск таймера
	fmt.Printf("sleeping for %d seconds\n", duration)
	<-timer.C //остановка таймера
	fmt.Println("waking up")
}

func sleep3(duration time.Duration) {
	fmt.Printf("sleeping for %d seconds\n", duration)
	<-time.After(duration * time.Second) //после duration секунд в канал поступают данные
	fmt.Println("waking up")
}

func main() {
	sleep1(5)
	sleep2(5)
	sleep3(5)
}
