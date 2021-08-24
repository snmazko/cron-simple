package cron

import (
	"strings"
	"time"
)

type weekPart struct {
	part
}

func (r *weekPart) toRun(t *time.Time) bool {
	value := int(t.Weekday())

	if r == nil {
		return true
	}
	if r.value != nil {
		return *r.value == value
	}

	return r.toRunWithoutValue(value)
}

func (r *weekPart) parse(str string) (*weekPart, error) {
	if p, err := parsePart(r.prepare(str), 0, 6); err != nil {
		return nil, err
	} else if p != nil {
		return &weekPart{part: *p}, nil
	}

	return nil, nil
}

func (r *weekPart) prepare(str string) string {
	return strings.Replace(str, "7", "0", -1)
}
