package main

import (
	"fmt"
	"github.com/natealcedo/intro-golang/task"
	"github.com/natealcedo/intro-golang/worker"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	numberOfWorkers := 1000
	numberOfTasks := 10000

	concreteWorkers := []worker.Worker{}

	// Initialize Workers
	for i := 0; i < numberOfWorkers; i++ {
		concreteWorker := worker.NewConcreteWorker(i, &wg)
		concreteWorkers[i] = concreteWorker
		concreteWorker.Start()
	}

	// Create and distribute tasks among concreteWorkers
	for i := 0; i < numberOfTasks; i++ {
		wg.Add(1)
		concreteWorkers[i%numberOfWorkers].ProcessTask(task.NewTask(i))
	}

	wg.Wait()
	fmt.Println("All tasks completed")
}
