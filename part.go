package cron

type part struct {
	value  *int
	values map[int]bool
	ranges *[2]int // index 0 - min value, 1 - max value
	period *int    // repetition frequency, minimum value one, maximum - maximum for the measured period
}

func (r *part) toRunWithoutValue(val int) bool {
	if r.values != nil {
		return r.isInValues(val)
	}
	if r.period != nil {
		return r.isInPeriod(val)
	} else if r.ranges != nil {
		return r.isInRanges(val)
	}

	return true
}

func (r *part) isInValues(val int) bool {
	return r.values[val]
}

func (r *part) isInRanges(val int) bool {
	return val >= r.ranges[0] && val <= r.ranges[1]
}

func (r *part) isInPeriod(val int) bool {
	if r.ranges == nil {
		return val%*r.period == 0
	}

	return r.isInRanges(val) && (val-r.ranges[0])%*r.period == 0
}
