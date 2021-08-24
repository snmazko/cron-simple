package cron

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RunnerMinute_toRun(t *testing.T) {
	at := assert.New(t)

	var r *minutePart
	at.True(r.toRun(toDate("2021-01-01T00:00:00Z")))

	r = &minutePart{part{}}
	at.True(r.toRun(toDate("2021-01-01T00:00:00Z")))

	r, err := r.parse("0")
	at.NoError(err)
	at.True(r.toRun(toDate("2021-01-01T00:00:00Z")))
	at.False(r.toRun(toDate("2021-01-01T00:01:00Z")))
	at.False(r.toRun(toDate("2020-01-01T00:59:00Z")))

	r, err = r.parse("59")
	at.NoError(err)
	at.True(r.toRun(toDate("2021-01-31T00:59:00Z")))
	at.False(r.toRun(toDate("2021-02-28T00:00:00Z")))
	at.False(r.toRun(toDate("2021-02-28T00:30:00Z")))
}
