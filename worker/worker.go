package worker

import "github.com/natealcedo/intro-golang/task"

type Worker interface {
	Start()
	ProcessTask(task *task.Task)
}
