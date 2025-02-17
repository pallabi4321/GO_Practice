package main

import (
	"fmt"
	"sync"
)

func generate(n int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- n
}

func square(n int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- n * n
}

func filterOutOdd(n int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	if n%2 == 0 {
		ch <- n
	}
}

func main() {
	totalNums := 100
	var wg1, wg2, wg3 sync.WaitGroup

	ch1, ch2, ch3 := make(chan int, totalNums), make(chan int, totalNums), make(chan int, totalNums)

	for i := 1; i <= totalNums; i++ {
		wg1.Add(1)
		go generate(i, ch1, &wg1)
	}

	wg1.Wait()
	close(ch1)

	for val := range ch1 {
		wg2.Add(1)
		go square(val, ch2, &wg2)
	}

	wg2.Wait()
	close(ch2)

	for val := range ch2 {
		wg3.Add(1)
		go filterOutOdd(val, ch3, &wg3)
	}

	wg3.Wait()
	close(ch3)

	for val := range ch3 {
		fmt.Println(val)
	}
}
