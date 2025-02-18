package main

import (
	"fmt"
	"sync"
)

func printEven(ech, och chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			<-ech
			fmt.Println("Even:", i)
			if i != 10 {
				och <- true
			}
		}
	}
}

func printOdd(ech, och chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			<-och
			fmt.Println("Odd:", i)
			ech <- true
		}
	}
}

func main() {
	ech := make(chan bool)
	och := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(2)
	go printOdd(ech, och, &wg)
	go printEven(ech, och, &wg)

	och <- true

	wg.Wait()
}
