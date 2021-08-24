package cron

import "time"

func intPtr(value int) *int {
	return &value
}

// date in format: 2020-11-01T00:00:00Z
func toDate(date string) *time.Time {
	res, err := time.Parse(time.RFC3339, date)
	if err != nil {
		println("Error parse time: " + err.Error())
	}

	return &res
}

func sleepMs(ms time.Duration) {
	time.Sleep(ms * time.Millisecond)
}

func dateWithLastDayOfMonth(t time.Time, addMonths ...int) time.Time {
	date := dateWithFirstDayOfMonth(t, addMonths...)

	return time.Date(
		date.Year(),
		date.Month(),
		dayInDateMonth(date),
		date.Hour(),
		date.Minute(),
		date.Second(),
		date.Nanosecond(),
		time.UTC,
	)
}

func dateWithFirstDayOfMonth(t time.Time, addMonths ...int) time.Time {
	date := time.Date(
		t.Year(),
		t.Month(),
		1,
		t.Hour(),
		t.Minute(),
		t.Second(),
		t.Nanosecond(),
		time.UTC,
	)

	if len(addMonths) > 0 {
		date = date.AddDate(0, addMonths[0], 0)
	}
	return date
}

func dayInDateMonth(t time.Time) int {
	switch t.Month() {
	case time.January:
		return 31
	case time.February:
		if isLeapYear(t.Year()) {
			return 29
		}
		return 28
	case time.March:
		return 31
	case time.April:
		return 30
	case time.May:
		return 31
	case time.June:
		return 30
	case time.July:
		return 31
	case time.August:
		return 31
	case time.September:
		return 30
	case time.October:
		return 31
	case time.November:
		return 30
	case time.December:
		return 31
	}
	return 0
}

// isLeapYear function time.isLeap from Go
func isLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
