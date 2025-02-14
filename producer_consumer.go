package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup

	for i := 1; i <= 13; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- rand.Intn(100)
		}()
	}

	go func() {
		for val := range ch {
			fmt.Println(val)
		}
	}()

	wg.Wait()
}
