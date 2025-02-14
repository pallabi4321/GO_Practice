package main

import "fmt"

func main() {
	ch := make(chan int, 5)

	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5

	close(ch)

	for val := range ch {
		fmt.Println(val)
	}
}
