package cron

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var allMinutes = ""
var tenMinute = ""
var oneHour = ""
var tenthDay = ""
var lastDay = ""
var firstMonth = ""
var sunday = ""

func clearVars() {
	allMinutes = ""
	tenMinute = ""
	oneHour = ""
	tenthDay = ""
	lastDay = ""
	firstMonth = ""
	sunday = ""
}

func Test_worker(t *testing.T) {
	at := assert.New(t)
	w := initWorker()
	createTasks(at, w)

	w.handleTasks(*toDate("2021-02-01T00:00:00Z"))
	sleepMs(10)
	at.Equal("run", allMinutes)
	at.Equal("", tenMinute)
	at.Equal("", oneHour)
	at.Equal("", tenthDay)
	at.Equal("", lastDay)
	at.Equal("", firstMonth)
	at.Equal("", sunday)
	clearVars()

	w.handleTasks(*toDate("2021-01-10T01:00:00Z"))
	sleepMs(10)
	at.Equal("run", allMinutes)
	at.Equal("", tenMinute)
	at.Equal("run", oneHour)
	at.Equal("run", tenthDay)
	at.Equal("", lastDay)
	at.Equal("run", firstMonth)
	at.Equal("run", sunday)
	clearVars()

	w.handleTasks(*toDate("2021-01-31T01:10:00Z"))
	sleepMs(10)
	at.Equal("run", allMinutes)
	at.Equal("run", tenMinute)
	at.Equal("run", oneHour)
	at.Equal("", tenthDay)
	at.Equal("run", lastDay)
	at.Equal("run", firstMonth)
	at.Equal("run", sunday)
	clearVars()
}

func createTasks(at *assert.Assertions, w Worker) {
	err := w.CreateTask("* * * * *", func() error {
		allMinutes = "run"
		return nil
	})
	at.NoError(err)

	err = w.CreateTask("10 * * * *", func() error {
		tenMinute = "run"
		return nil
	})
	at.NoError(err)

	err = w.CreateTask("* 1 * * *", func() error {
		oneHour = "run"
		return nil
	})
	at.NoError(err)

	err = w.CreateTask("* * 10 * *", func() error {
		tenthDay = "run"
		return nil
	})
	at.NoError(err)

	err = w.CreateTask("* * 0 * *", func() error {
		lastDay = "run"
		return nil
	})
	at.NoError(err)

	err = w.CreateTask("* * * 1 *", func() error {
		firstMonth = "run"
		return nil
	})
	at.NoError(err)

	err = w.CreateTask("* * * * 0", func() error {
		sunday = "run"
		return nil
	})
}
