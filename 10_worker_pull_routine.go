// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Job struct {
	ID       int
	RandomNo int
}

type Result struct {
	JobDetail      Job
	SumOfRandomNos int
}

func main() {
	chj := make(chan Job, 100)
	chr := make(chan Result, 100)

	var wgj sync.WaitGroup
	var wgr sync.WaitGroup

	for i := 1; i <= 100; i++ {
		wgj.Add(1)
		go func(id int) {
			defer wgj.Done()
			chj <- Job{ID: i, RandomNo: rand.Intn(1000)}
		}(i)
	}

	go func() {
		wgj.Wait()
		close(chj)
	}()

	for i := 0; i < 10; i++ {
		wgr.Add(1)
		go func() {
			wgr.Done()
			for job := range chj {
				sum := 0
				rno := job.RandomNo
				for rno > 0 {
					sum += rno % 10
					rno = rno / 10
				}
				chr <- Result{JobDetail: job, SumOfRandomNos: sum}
			}
		}()
	}

	go func() {
		wgr.Wait()
		close(chr)
	}()

	for v := range chr {
		fmt.Println(v)
	}
}
