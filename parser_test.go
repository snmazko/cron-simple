package cron

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_parse(t *testing.T) {
	at := assert.New(t)

	_, err := parse("")
	at.Error(err)

	_, err = parse("* * *")
	at.Error(err)

	_, err = parse("* * * *")
	at.Error(err)

	_, err = parse("* * * * * *")
	at.Error(err)

	ss, err := parse("   */10    *   0   *  7")
	at.NoError(err)
	at.Equal([]string{"*/10", "*", "0", "*", "7"}, ss)
}

// 9
func Test_parsePartByValue(t *testing.T) {
	at := assert.New(t)

	_, err := parsePartByValue("", 1, 12)
	at.Error(err)

	_, err = parsePartByValue("0", 1, 12)
	at.Error(err)

	_, err = parsePartByValue("13", 1, 12)
	at.Error(err)

	res, err := parsePartByValue("0", 0, 12)
	at.NoError(err)
	at.Equal(&part{
		value: intPtr(0),
	}, res)

	res, err = parsePartByValue("6", 1, 12)
	at.NoError(err)
	at.Equal(&part{
		value: intPtr(6),
	}, res)
}

// 1,6,8
func Test_parsePartByMultiValue(t *testing.T) {
	at := assert.New(t)

	_, err := parsePartByMultiValue("", 1, 12)
	at.Error(err)

	_, err = parsePartByMultiValue("0", 1, 12)
	at.Error(err)

	_, err = parsePartByMultiValue("1,8,13", 1, 12)
	at.Error(err)

	res, err := parsePartByMultiValue("1,6,8", 1, 12)
	at.NoError(err)
	at.Len(res.values, 3)
	at.Equal(true, res.values[1])
	at.Equal(true, res.values[6])
	at.Equal(true, res.values[8])

	res.values = nil
	at.Equal(&part{}, res)
}

// 3-5
func Test_parsePartByRange(t *testing.T) {
	at := assert.New(t)

	_, err := parsePartByRange("", 2, 12)
	at.Error(err)

	_, err = parsePartByRange("0", 2, 12)
	at.Error(err)

	_, err = parsePartByRange("1-3", 2, 12)
	at.Error(err)

	res, err := parsePartByRange("3-5", 2, 12)
	at.NoError(err)
	at.Equal(&part{
		ranges: &[2]int{3, 5},
	}, res)
}

// */2	1-6/2
func Test_parsePartByPeriod_error(t *testing.T) {
	at := assert.New(t)

	// out of bounds in the first argument
	_, err := parsePartByPeriod("1/2", 2, 12)
	at.Error(err)
	_, err = parsePartByPeriod("13/2", 2, 12)
	at.Error(err)
	_, err = parsePartByPeriod("1-6/2", 2, 12)
	at.Error(err)

	// out of bounds in the second argument
	_, err = parsePartByPeriod("*/0", 2, 12)
	at.Error(err)
	_, err = parsePartByPeriod("*/13", 2, 12)
	at.Error(err)
	_, err = parsePartByPeriod("2-6/0", 2, 12)
	at.Error(err)
}

// */2	1-6/2
func Test_parsePartByPeriod_success(t *testing.T) {
	at := assert.New(t)

	res, err := parsePartByPeriod("*/1", 2, 12)
	at.NoError(err)
	at.Equal(&part{
		period: intPtr(1),
	}, res)

	res, err = parsePartByPeriod("3-5/2", 2, 12)
	at.NoError(err)
	at.Equal(&part{
		period: intPtr(2),
		ranges: &[2]int{3, 5},
	}, res)
}
