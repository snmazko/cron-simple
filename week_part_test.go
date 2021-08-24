package cron

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_weekPart_toRun(t *testing.T) {
	at := assert.New(t)

	var r *weekPart
	at.True(r.toRun(toDate("2021-01-01T00:00:00Z")))

	r = &weekPart{part{}}
	at.True(r.toRun(toDate("2021-01-01T00:00:00Z")))

	r, err := r.parse("*")
	at.NoError(err)
	at.True(r.toRun(toDate("2021-01-01T00:00:00Z")))

	r, err = r.parse("0")
	at.NoError(err)
	at.True(r.toRun(toDate("2021-01-24T00:00:00Z")))
	at.False(r.toRun(toDate("2021-01-23T00:01:00Z")))
	at.False(r.toRun(toDate("2020-01-25T00:59:00Z")))

	r, err = r.parse("7")
	at.NoError(err)
	at.True(r.toRun(toDate("2021-01-24T00:00:00Z")))
	at.False(r.toRun(toDate("2021-01-23T00:01:00Z")))
	at.False(r.toRun(toDate("2020-01-25T00:59:00Z")))

	r, err = r.parse("0,7")
	at.NoError(err)
	at.True(r.toRun(toDate("2021-01-24T00:00:00Z")))
	at.False(r.toRun(toDate("2021-01-23T00:01:00Z")))
	at.False(r.toRun(toDate("2020-01-25T00:59:00Z")))
}
