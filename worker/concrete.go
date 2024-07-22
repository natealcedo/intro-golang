package worker

import (
	"fmt"
	"github.com/natealcedo/intro-golang/task"
	"math/rand"
	"sync"
	"time"
)

type ConcreteWorker struct {
	id   int
	task chan *task.Task
	wg   *sync.WaitGroup
}

func (w *ConcreteWorker) Start() {
	go func() {
		for concreteTask := range w.task {
			fmt.Printf("Task %d processing concreteTask: %+v\n", w.id, *concreteTask)
			// multiply this by a random number between 1 - 5
			time.Sleep(time.Second * time.Duration(rand.Intn(5)))
			w.wg.Done() // Signal concreteTask completion here, inside the worker's goroutine
		}
	}()
}

func (w *ConcreteWorker) ProcessTask(task *task.Task) {
	w.task <- task
}

func NewConcreteWorker(id int, wg *sync.WaitGroup) *ConcreteWorker {
	return &ConcreteWorker{
		id:   id,
		task: make(chan *task.Task),
		wg:   wg,
	}
}
