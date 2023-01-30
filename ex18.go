package main

import (
	"fmt"
	"sync"
)

type counter struct {
	val int
	mut sync.Mutex
}

func (c *counter) add(n int, wg *sync.WaitGroup) {
	c.mut.Lock()   //блокируем для других горутин
	c.val += n     //добавляем
	c.mut.Unlock() //разблокируем
}

func main() {
	var c counter
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 1000; j++ {
				c.add(1, &wg)
			}
			wg.Done()
		}()
	}
	wg.Wait() //ожидаем завершения горутин
	fmt.Println(c.val)
}
