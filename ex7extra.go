package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func storeVal(i int, m *sync.Map, wg *sync.WaitGroup) {
	for j := 0; j < 100; j++ {
		m.Store(100*i+j, rand.Intn(100)) //добалвение
	}
	wg.Done()
}

func main() {
	var m = sync.Map{} //специальная map для конкурентной записи
	var wg sync.WaitGroup
	wg.Add(4)
	for i := 0; i < 4; i++ {
		go storeVal(i, &m, &wg)
	}
	wg.Wait()
	amount := 0
	m.Range(func(key, value any) bool {
		fmt.Printf("%d: %d   ", key, value) //печать ключа и значения
		amount++
		return true
	})
	fmt.Println(amount)
}
