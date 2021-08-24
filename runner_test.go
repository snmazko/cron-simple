package cron

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CreateCron(t *testing.T) {
	at := assert.New(t)

	res, err := CreateRunner("")
	at.Error(err)

	res, err = CreateRunner("* * * * *")
	at.NoError(err)
	at.Equal(&runner{}, res)

	res, err = CreateRunner("30 10-40/2 * 1,3,5 7")
	at.NoError(err)

	r := res.(*runner)
	at.Len(r.month.values, 3)
	at.Equal(true, r.month.values[1])
	at.Equal(true, r.month.values[3])
	at.Equal(true, r.month.values[5])
	r.month.values = nil

	at.Equal(&runner{
		week: &weekPart{part{
			value: intPtr(0),
		}},

		month: &monthPart{
			part: part{},
		},

		hour: &hourPart{
			part: part{
				ranges: &[2]int{10, 40},
				period: intPtr(2),
			},
		},

		minute: &minutePart{
			part: part{
				value: intPtr(30),
			},
		},
	}, res)
}

func Test_Cron_ToRun(t *testing.T) {
	at := assert.New(t)

	res, err := CreateRunner("30 10-20/2 * 1,3,5 *")
	at.NoError(err)
	at.True(res.ToRun(*toDate("2021-01-01T10:30:00Z")))
	at.True(res.ToRun(*toDate("2021-03-21T16:30:00Z")))
	at.True(res.ToRun(*toDate("2021-05-25T20:30:00Z")))
	// wrong minutes - 31
	at.False(res.ToRun(*toDate("2021-01-01T10:31:00Z")))
	// wrong hour - 08
	at.False(res.ToRun(*toDate("2021-01-01T08:30:00Z")))
	// wrong minutes - 11, fit only 10, 12, 14, 16, 18, 20
	at.False(res.ToRun(*toDate("2021-01-01T11:30:00Z")))
	// wrong month - 02, fit only 1,3,5
	at.False(res.ToRun(*toDate("2021-02-01T10:30:00Z")))

	res, err = CreateRunner("30 10-20/2 * 1,3,5 5")
	at.NoError(err)
	at.True(res.ToRun(*toDate("2021-01-01T10:30:00Z")))
	at.True(res.ToRun(*toDate("2021-03-12T16:30:00Z")))
	// wrong week. May 25 is Tuesday(2), and it should work on Friday(5)
	at.False(res.ToRun(*toDate("2021-05-25T20:30:00Z")))
}
