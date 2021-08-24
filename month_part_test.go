package cron

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RunnerMonth_toRun(t *testing.T) {
	at := assert.New(t)

	var r *monthPart
	at.True(r.toRun(toDate("2021-01-01T00:00:00Z")))

	r = &monthPart{part{}}
	at.True(r.toRun(toDate("2021-01-01T00:00:00Z")))

	r, err := r.parse("*")
	at.NoError(err)
	at.True(r.toRun(toDate("2021-01-01T00:00:00Z")))

	r, err = r.parse("1")
	at.NoError(err)
	at.True(r.toRun(toDate("2021-01-01T00:00:00Z")))
	at.False(r.toRun(toDate("2021-02-01T00:00:00Z")))
}

func Test_RunnerMonth_toRunWithoutValue(t *testing.T) {
	at := assert.New(t)

	var r *monthPart

	r, err := r.parse("1,12")
	at.NoError(err)
	at.True(r.toRun(toDate("2021-01-01T00:00:00Z")))
	at.True(r.toRun(toDate("2021-12-01T00:00:00Z")))
	at.False(r.toRun(toDate("2021-06-01T00:00:00Z")))

	r, err = r.parse("3-5")
	at.NoError(err)
	at.True(r.toRun(toDate("2021-03-01T00:00:00Z")))
	at.True(r.toRun(toDate("2021-05-01T00:00:00Z")))
	at.False(r.toRun(toDate("2021-06-01T00:00:00Z")))

	r, err = r.parse("3-9/2")
	at.NoError(err)
	at.False(r.toRun(toDate("2021-02-01T00:00:00Z")))
	at.False(r.toRun(toDate("2021-04-01T00:00:00Z")))
	at.False(r.toRun(toDate("2021-06-01T00:00:00Z")))
	at.False(r.toRun(toDate("2021-08-01T00:00:00Z")))
	at.False(r.toRun(toDate("2021-10-01T00:00:00Z")))
	at.True(r.toRun(toDate("2021-03-01T00:00:00Z")))
	at.True(r.toRun(toDate("2021-05-01T00:00:00Z")))
	at.True(r.toRun(toDate("2021-07-01T00:00:00Z")))
	at.True(r.toRun(toDate("2021-09-01T00:00:00Z")))

	r, err = r.parse("*/2")
	at.NoError(err)
	at.True(r.toRun(toDate("2021-02-01T00:00:00Z")))
	at.True(r.toRun(toDate("2021-04-01T00:00:00Z")))
	at.True(r.toRun(toDate("2021-06-01T00:00:00Z")))
	at.True(r.toRun(toDate("2021-08-01T00:00:00Z")))
	at.True(r.toRun(toDate("2021-10-01T00:00:00Z")))
	at.False(r.toRun(toDate("2021-03-01T00:00:00Z")))
	at.False(r.toRun(toDate("2021-05-01T00:00:00Z")))
	at.False(r.toRun(toDate("2021-07-01T00:00:00Z")))
	at.False(r.toRun(toDate("2021-09-01T00:00:00Z")))

}
