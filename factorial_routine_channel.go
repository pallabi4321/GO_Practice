package main

import (
	"fmt"
	"sync"
)

func factorial(n int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fact := 1
	for i := 1; i <= n; i++ {
		fact *= i
	}
	ch <- fact
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	for i := 3; i <= 8; i++ {
		wg.Add(1)
		go factorial(i, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for val := range ch {
		fmt.Println(val)
	}

}
