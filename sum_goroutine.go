// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
)

func sum(n1 int, n2 int, ch chan int, wr *sync.WaitGroup) {
	wr.Done()
	ch <- n1 + n2
}

func main() {
	var wr sync.WaitGroup
	ch := make(chan int)

	wr.Add(1)
	go sum(5, 6, ch, &wr)

	wr.Wait()
	fmt.Println("sum: ", <-ch)

}