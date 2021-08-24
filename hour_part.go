package cron

import (
	"time"
)

type hourPart struct {
	part
}

func (r *hourPart) toRun(t *time.Time) bool {
	value := t.Hour()

	if r == nil {
		return true
	}
	if r.value != nil {
		return *r.value == value
	}

	return r.toRunWithoutValue(value)
}

func (r *hourPart) parse(str string) (*hourPart, error) {
	if p, err := parsePart(str, 0, 59); err != nil {
		return nil, err
	} else if p != nil {
		return &hourPart{part: *p}, nil
	}

	return nil, nil
}
