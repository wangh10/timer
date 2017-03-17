package timer

type Task interface {
	Do()
}

type task struct {
	t Task
	c uint
}
