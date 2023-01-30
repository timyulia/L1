package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func stopContext(ctx context.Context, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done(): //из контекста сигнал завершения
			fmt.Println("3: stopped")
			wg.Done()
			return
		default:
			fmt.Println("3: working")
			time.Sleep(time.Second)
		}
	}
}

func stopClosed(ch1 chan struct{}, wg *sync.WaitGroup) {
	for {
		select {
		case _, opened := <-ch1:

			if !opened { //канал закрыт
				fmt.Println("1: stopped")
				wg.Done()
				return
			}
		default:
			fmt.Println("1: working")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func stopChan(ch chan bool, wg *sync.WaitGroup) {
	for {
		select {
		case <-ch: //получено значение из канала
			fmt.Println("2: stopped")
			wg.Done()
			return
		default:
			fmt.Println("2: working")
			time.Sleep(250 * time.Millisecond)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(4)
	ch1 := make(chan struct{})
	ch2 := make(chan bool)
	ctxt, cancel := context.WithCancel(context.Background()) //создание контекста, чтобы передать его параметром

	go stopContext(ctxt, &wg) //остановка контекстом
	go stopClosed(ch1, &wg)   //остановка закрытием канала
	go stopChan(ch2, &wg)     //остановка каналом

	flag := true
	go func() {
		for flag { //изменение переменной флаг в main будет останавливать горутину
			fmt.Println("4: working")
			time.Sleep(400 * time.Millisecond)
		}
		fmt.Println("4: stopped")
		wg.Done()
		return
	}()

	time.Sleep(3 * time.Second)
	close(ch1) //закрытие канала

	time.Sleep(3 * time.Second)
	ch2 <- true //передача в канал значения, чтобы остановить горутину

	time.Sleep(3 * time.Second)
	cancel() //закрытие контекстом

	time.Sleep(3 * time.Second)
	flag = false //изменение флага на false, чтобы остановить горутину

	wg.Wait()
}
