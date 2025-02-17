package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker function that processes tasks
func worker(id int, tasks <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Worker %d: Processing task %d\n", id, task)
		time.Sleep(time.Second) // Simulating work
	}
	fmt.Printf("Worker %d: Exiting\n", id)
}

func main() {
	// Define the number of workers and tasks
	const numWorkers = 5
	const numTasks = 10

	// Create a channel for tasks
	tasks := make(chan int, numTasks)
	var wg sync.WaitGroup

	// Send tasks to the channel
	for i := 1; i <= numTasks; i++ {
		tasks <- i
	}

	// Start worker goroutines
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, &wg)
	}

	// Close the tasks channel (all tasks sent)
	close(tasks)

	// Wait for all workers to complete
	wg.Wait()

	fmt.Println("All tasks completed!")
}
