package cron

import (
	"regexp"
	"time"
)

var multiSpace = regexp.MustCompile("\\s+")

// min	h	day	mon	week
// *	*	*	*	*	- every minute
// */5	*	*	*	*	- every 5 minutes
// 0	*	*	*	*	- every hour
// 0	*/3	*	*	*	- every 3 hours
// 0	13	*	*	*	- every day at 13:00
// 30	2	*	*	*	- every day at 2:30
// 0	0	*	*	*	- every day at midnight
// 0	0	1	*	*	- on the first day of every month
// 0	0	0	*	*	- on the last day of every month
// 0	0	6	1	*	- on the sixth day of the first month
// 0	0	*	*	0	- at midnight every Sunday
type runner struct {
	minute *minutePart
	hour   *hourPart
	day    *dayPart
	month  *monthPart
	week   *weekPart
}

func (p *runner) ToRun(t time.Time) bool {
	isRun := p.week.toRun(&t) &&
		p.month.toRun(&t) &&
		p.day.toRun(&t) &&
		p.hour.toRun(&t) &&
		p.minute.toRun(&t)

	return isRun
}

func CreateRunner(str string) (Runnable, error) {
	r := &runner{}
	ss, err := parse(str)
	if err != nil {
		return nil, err
	}
	if r.minute, err = r.minute.parse(ss[0]); err != nil {
		return nil, err
	}
	if r.hour, err = r.hour.parse(ss[1]); err != nil {
		return nil, err
	}
	if r.day, err = r.day.parse(ss[2]); err != nil {
		return nil, err
	}
	if r.month, err = r.month.parse(ss[3]); err != nil {
		return nil, err
	}
	if r.week, err = r.week.parse(ss[4]); err != nil {
		return nil, err
	}

	return r, nil
}
