package cron

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const layoutParseDate = "2006-01-02"

func Test_DateWithFirstDayOfMonth(t *testing.T) {
	var at = assert.New(t)
	expect, _ := time.Parse(layoutParseDate, "2022-11-01")

	date, _ := time.Parse(layoutParseDate, "2022-10-01")
	at.Equal(expect, dateWithFirstDayOfMonth(date, 1))

	date, _ = time.Parse(layoutParseDate, "2022-10-15")
	at.Equal(expect, dateWithFirstDayOfMonth(date, 1))

	date, _ = time.Parse(layoutParseDate, "2022-10-31")
	at.Equal(expect, dateWithFirstDayOfMonth(date, 1))
}

func Test_DateWithLastDayOfMonth(t *testing.T) {
	var at = assert.New(t)
	expect, _ := time.Parse(layoutParseDate, "2022-11-30")

	date, _ := time.Parse(layoutParseDate, "2022-10-01")
	at.Equal(expect, dateWithLastDayOfMonth(date, 1))

	date, _ = time.Parse(layoutParseDate, "2022-10-15")
	at.Equal(expect, dateWithLastDayOfMonth(date, 1))

	date, _ = time.Parse(layoutParseDate, "2022-10-31")
	at.Equal(expect, dateWithLastDayOfMonth(date, 1))

	//not Leap Year
	expect, _ = time.Parse(layoutParseDate, "2022-02-28")
	date, _ = time.Parse(layoutParseDate, "2022-02-01")
	at.Equal(expect, dateWithLastDayOfMonth(date))

	//Is Leap Year
	expect, _ = time.Parse(layoutParseDate, "2020-02-29")
	date, _ = time.Parse(layoutParseDate, "2020-02-01")
	at.Equal(expect, dateWithLastDayOfMonth(date))
}
