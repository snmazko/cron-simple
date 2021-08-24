package cron_test

import (
	"fmt"
	"github.com/snmazko/cron-simple"
	"testing"
	"time"
)

type logger struct{}

func (l logger) Errorf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

func (l logger) Info(data interface{}) {
	fmt.Printf("%v\n", data)
}

func TestExample(t *testing.T) {
	worker := cron.NewWorker().
		SetLogger(logger{}).
		Start().
		Start() // Double call for no block test only

	worker.MustCreateTask("* * * * *", func1)           //every minute
	worker.MustCreateTask("*/2 * * * *", func() error { //every two minute
		return func2("Run func2")
	})

	//time.Sleep(5 * time.Minute)

	worker.Stop()
	worker.Stop() // Double call for no block test only
}

func func1() error {
	fmt.Printf("Run func1, time: %v\n", time.Now())
	return nil
}

func func2(msg string) error {
	fmt.Printf("%s, time: %v\n", msg, time.Now())
	return nil
}
