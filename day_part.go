package cron

import (
	"time"
)

type dayPart struct {
	part
}

func (r *dayPart) toRun(t *time.Time) bool {
	value := t.Day()

	if r == nil {
		return true
	}
	if r.value != nil {
		if *r.value > 0 {
			return *r.value == value
		}
		return dateWithLastDayOfMonth(*t).Day() == value
	}

	return r.toRunWithoutValue(value)
}

func (r *dayPart) parse(str string) (*dayPart, error) {
	if p, err := parsePart(str, 0, 31); err != nil {
		return nil, err
	} else if p != nil {
		return &dayPart{part: *p}, nil
	}

	return nil, nil
}
