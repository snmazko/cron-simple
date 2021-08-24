package cron

import (
	"time"
)

type minutePart struct {
	part
}

func (r *minutePart) toRun(t *time.Time) bool {
	value := t.Minute()

	if r == nil {
		return true
	}
	if r.value != nil {
		return *r.value == value
	}

	return r.toRunWithoutValue(value)
}

func (r *minutePart) parse(str string) (*minutePart, error) {
	if p, err := parsePart(str, 0, 59); err != nil {
		return nil, err
	} else if p != nil {
		return &minutePart{part: *p}, nil
	}

	return nil, nil
}
