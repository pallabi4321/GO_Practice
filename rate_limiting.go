package main

import (
	"fmt"
	"sync"
	"time"
)

func rateLimit(limit int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	tick := time.NewTicker(time.Second * 3 / time.Duration(limit))
	defer tick.Stop()
	for val := range ch {
		<-tick.C
		fmt.Println(val)
	}
}

func main() {
	limit := 3
	var wg sync.WaitGroup
	ch := make(chan int, 10)

	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)

	wg.Add(1)
	go rateLimit(limit, ch, &wg)

	wg.Wait()
}
