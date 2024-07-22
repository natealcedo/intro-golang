package task

type Task struct {
	id int
}

func NewTask(id int) *Task {
	return &Task{id: id}
}
