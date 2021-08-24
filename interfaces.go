package cron

import (
	"time"
)

type Logger interface {
	Errorf(format string, a ...interface{})
	Info(data interface{})
}

type Runnable interface {
	ToRun(time.Time) bool
}

type Executable interface {
	Handle() error
}

type Task interface {
	Runnable
	Executable
}

type Worker interface {
	Start() Worker
	Stop()
	MustCreateTask(str string, handler func() error) Worker
	CreateTask(str string, handler func() error) error
	AddTask(task Task) Worker
	SetLogger(Logger) Worker
}
