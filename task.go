package cron

import "time"

type task struct {
	runner  Runnable
	handler func() error
}

func NewTask(runner Runnable, handler func() error) Task {
	return &task{runner: runner, handler: handler}
}

func CreateTask(str string, handler func() error) (Task, error) {
	if r, err := CreateRunner(str); err != nil {
		return nil, err
	} else {
		return NewTask(r, handler), err
	}
}

func (t *task) ToRun(time time.Time) bool {
	return t.runner.ToRun(time)
}

func (t *task) Handle() error {
	return t.handler()
}
