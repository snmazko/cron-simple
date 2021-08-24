package cron

import (
	"fmt"
	"sync"
	"time"
)

type worker struct {
	logger         Logger
	tasks          []Task
	isStop         bool
	stop           chan struct{}
	runningTasks   map[Task]bool
	runningTasksMx sync.Mutex
	startOnceMx    sync.Once
	stopOnceMx     sync.Once

	sync.Mutex
}

func NewWorker() Worker {
	return initWorker()
}

func initWorker() *worker {
	return &worker{
		stop:         make(chan struct{}),
		runningTasks: make(map[Task]bool),
	}
}

func (rec *worker) Start() Worker {
	rec.startOnceMx.Do(func() {
		go rec.run()
		go rec.waitStop()
	})

	return rec
}

func (rec *worker) Stop() {
	rec.stopOnceMx.Do(func() {
		rec.stop <- struct{}{}
	})
}

func (rec *worker) MustCreateTask(str string, handler func() error) Worker {
	err := rec.CreateTask(str, handler)
	if err != nil {
		panic(err)
	}

	return rec
}

func (rec *worker) CreateTask(str string, handler func() error) error {
	t, err := CreateTask(str, handler)
	if err != nil {
		return err
	}

	rec.AddTask(t)

	return nil
}

func (rec *worker) AddTask(task Task) Worker {
	rec.Lock()
	defer rec.Unlock()

	rec.tasks = append(rec.tasks, task)

	return rec
}

func (rec *worker) SetLogger(log Logger) Worker {
	rec.logger = log

	return rec
}

func (rec *worker) run() {
	rec.log("Cron task manager started")
	for {
		if rec.isStop {
			break
		}
		rec.handleTasks(time.Now())
		rec.wait()
	}
	rec.log("Cron task manager stopped")
}

func (rec *worker) handleTasks(t time.Time) {
	for _, w := range rec.getTasks() {
		if w.ToRun(t) {
			go rec.runTask(w)
		}
	}
}

func (rec *worker) getTasks() []Task {
	rec.Lock()
	defer rec.Unlock()

	w := make([]Task, len(rec.tasks))
	copy(w, rec.tasks)

	return w
}

func (rec *worker) wait() {
	time.Sleep(rec.delay())
}

func (rec *worker) delay() time.Duration {
	return time.Duration(60-time.Now().Second()+1) * time.Second
}

func (rec *worker) waitStop() {
	<-rec.stop
	rec.isStop = true
}

func (rec *worker) runTask(w Task) {
	if rec.isRunning(w) {
		return
	}
	rec.log(fmt.Sprintf("Task started: %T", w))

	defer func() {
		rec.log(fmt.Sprintf("Finished task: %T", w))
		if err := recover(); err != nil {
			rec.error(
				"Panic recovery for task:\nTask: %T\nError: %+v", w, err,
			)
		}
	}()

	rec.setIsRunning(w, true)
	defer func() {
		rec.setIsRunning(w, false)
	}()

	if err := w.Handle(); err != nil {
		rec.error(
			"Error in task:\nTask: %T\nError: %+v", w, err,
		)
	}
}

func (rec *worker) isRunning(w Task) bool {
	rec.runningTasksMx.Lock()
	defer rec.runningTasksMx.Unlock()

	return rec.runningTasks[w]
}

func (rec *worker) setIsRunning(w Task, isRun bool) {
	rec.runningTasksMx.Lock()
	defer rec.runningTasksMx.Unlock()

	rec.runningTasks[w] = isRun
}

func (rec *worker) log(msg string) {
	if rec.logger != nil {
		rec.logger.Info(msg)
	}
}

func (rec *worker) error(format string, a ...interface{}) {
	if rec.logger != nil {
		rec.logger.Errorf(format, a...)
	}
}
