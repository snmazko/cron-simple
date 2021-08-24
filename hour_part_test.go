package cron

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RunnerHour_toRun(t *testing.T) {
	at := assert.New(t)

	var r *hourPart
	at.True(r.toRun(toDate("2021-01-01T00:00:00Z")))

	r = &hourPart{part{}}
	at.True(r.toRun(toDate("2021-01-01T00:00:00Z")))

	r, err := r.parse("0")
	at.NoError(err)
	at.True(r.toRun(toDate("2021-01-01T00:00:00Z")))
	at.False(r.toRun(toDate("2021-01-01T01:00:00Z")))
	at.False(r.toRun(toDate("2020-01-01T23:00:00Z")))

	r, err = r.parse("23")
	at.NoError(err)
	at.True(r.toRun(toDate("2021-01-31T23:00:00Z")))
	at.False(r.toRun(toDate("2021-02-28T00:00:00Z")))
	at.False(r.toRun(toDate("2021-02-28T12:00:00Z")))
}
