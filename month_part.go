package cron

import (
	"time"
)

type monthPart struct {
	part
}

func (r *monthPart) toRun(t *time.Time) bool {
	value := int(t.Month())

	if r == nil {
		return true
	}
	if r.value != nil {
		return *r.value == value
	}

	return r.toRunWithoutValue(value)
}

func (r *monthPart) parse(str string) (*monthPart, error) {
	if p, err := parsePart(str, 1, 12); err != nil {
		return nil, err
	} else if p != nil {
		return &monthPart{part: *p}, nil
	}

	return nil, nil
}
