package main

import (
	"fmt"
	"sync"
	"time"
)

type Worker interface {
	Start(wg *sync.WaitGroup)
}

type Work struct {
	id   int
	task chan int
}

func NewWork(id int) *Work {
	return &Work{
		id:   id,
		task: make(chan int),
	}
}

func (w *Work) Start(wg *sync.WaitGroup) {
	go func() {
		for task := range w.task {
			fmt.Printf("Work %d processing task: %d\n", w.id, task)
			time.Sleep(time.Second) // Simulate task processing time
			wg.Done()               // Signal task completion here, inside the worker's goroutine
		}
	}()
}

func (w *Work) ProcessTask(task int) {
	w.task <- task
}

func main() {
	var wg sync.WaitGroup
	workToBeDone := make([]*Work, 5)

	// Initialize workers and share the same WaitGroup among them
	for i := range workToBeDone {
		workToBeDone[i] = NewWork(i)
		wg.Add(1) // Increment here for each worker
		go workToBeDone[i].Start(&wg)
	}

	tasks := []int{1, 2, 3, 4, 5}
	taskChannel := make(chan int, len(tasks))

	// Send tasks to the channel
	go func() {
		for _, task := range tasks {
			taskChannel <- task
		}
		close(taskChannel) // Close the channel to signal no more tasks will be sent
	}()

	// Distribute tasks to workers
	for task := range taskChannel {
		workerIndex := task % len(workToBeDone)
		workToBeDone[workerIndex].ProcessTask(task)
	}

	// Close each worker's task channel to signal no more tasks will be sent
	for _, worker := range workToBeDone {
		close(worker.task)
	}

	wg.Wait() // Wait for all tasks to be processed
	fmt.Println("All tasks processed.")
}
