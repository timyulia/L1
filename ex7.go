package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func add(i int, m map[int]int, mut *sync.Mutex, wg *sync.WaitGroup) {
	for j := 0; j < 10; j++ {
		mut.Lock()                 //блокировка
		m[i*10+j] = rand.Intn(100) //добавление
		mut.Unlock()               //разблокировка
	}
	wg.Done()
}

func main() {
	n := 4
	var wg sync.WaitGroup
	var mut sync.Mutex
	m := make(map[int]int)
	wg.Add(n)
	for i := 0; i < n; i++ {
		go add(i, m, &mut, &wg)
	}
	wg.Wait() //ожидание выполнения горутин
	fmt.Println(m)
}
