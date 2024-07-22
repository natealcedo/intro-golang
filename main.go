package main

import (
	"fmt"
	"sync"
	"time"
)

type Worker interface {
	Start()
	ProcessTask(task Task)
}

// ConcreteWorker is a worker that processes tasks
type ConcreteWorker struct {
	id   int
	task chan Task
	wg   *sync.WaitGroup
}

func (w *ConcreteWorker) Start() {
	go func() {
		for task := range w.task {
			fmt.Printf("Task %d processing task: %d\n", w.id, task)
			time.Sleep(time.Second) // Simulate task processing time
			w.wg.Done()             // Signal task completion here, inside the worker's goroutine
		}
	}()
}

// Task struct
type Task struct {
	id int
}

func (w *ConcreteWorker) ProcessTask(task Task) {
	w.task <- task
}

func main() {
	wg := sync.WaitGroup{}
	numberOfWorkers := 1000
	numberOfTasks := 10000

	workers := make([]Worker, numberOfWorkers)

	// Initialize Workers
	for i := 0; i < numberOfWorkers; i++ {
		worker := &ConcreteWorker{
			id:   i,
			task: make(chan Task),
			wg:   &wg,
		}
		workers[i] = worker
		worker.Start()
	}

	// Create and distribute tasks among workers
	for i := 0; i < numberOfTasks; i++ {
		wg.Add(1)
		workers[i%numberOfWorkers].ProcessTask(Task{id: i})
	}

	wg.Wait()
	fmt.Println("All tasks completed")
}
