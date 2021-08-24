package cron

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RunnerDay_toRun(t *testing.T) {
	at := assert.New(t)

	var r *dayPart
	at.True(r.toRun(toDate("2021-01-01T00:00:00Z")))

	r = &dayPart{part{}}
	at.True(r.toRun(toDate("2021-01-01T00:00:00Z")))

	r, err := r.parse("1")
	at.NoError(err)
	at.True(r.toRun(toDate("2021-01-01T00:00:00Z")))
	at.False(r.toRun(toDate("2021-01-02T00:00:00Z")))
	at.False(r.toRun(toDate("2020-12-31T00:00:00Z")))

	r, err = r.parse("31")
	at.NoError(err)
	at.True(r.toRun(toDate("2021-01-31T00:00:00Z")))
	at.False(r.toRun(toDate("2021-02-28T00:00:00Z")))

	// last day of month
	r, err = r.parse("0")
	at.NoError(err)
	at.True(r.toRun(toDate("2021-01-31T00:00:00Z")))
	at.False(r.toRun(toDate("2021-01-30T00:00:00Z")))
	at.True(r.toRun(toDate("2021-02-28T00:00:00Z")))
	at.False(r.toRun(toDate("2021-02-27T00:00:00Z")))
}
