package cron

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	digitPat    = "\\d{1,2}"
	multiPat    = "\\d+,[\\d,]+"
	rangePat    = "\\d{1,2}-\\d{1,2}"
	fractionPat = "^(\\*|(" + rangePat + "))/" + digitPat + "$"
	countPart   = 5
)

var digit = regexp.MustCompile("^" + digitPat + "$")
var multiValues = regexp.MustCompile("^" + multiPat + "$")
var rangeValues = regexp.MustCompile("^" + rangePat + "$")
var periodValues = regexp.MustCompile(fractionPat)

func parse(str string) ([]string, error) {
	ss := strings.Split(strings.TrimSpace(multiSpace.ReplaceAllString(str, " ")), " ")
	if len(ss) != countPart {
		return nil, errors.New(
			fmt.Sprintf("The worker config string must contain %d elements separated by a space\nOrigin: %s", countPart, str),
		)
	}

	return ss, nil
}

func parsePart(str string, min, max int) (*part, error) {
	if str == "*" {
		return nil, nil
	}
	if digit.MatchString(str) {
		return parsePartByValue(str, min, max)
	}
	if multiValues.MatchString(str) {
		return parsePartByMultiValue(str, min, max)
	}
	if rangeValues.MatchString(str) {
		return parsePartByRange(str, min, max)
	}
	if periodValues.MatchString(str) {
		return parsePartByPeriod(str, min, max)
	}

	return nil, errors.New(
		fmt.Sprintf("Invalid value format: %s", str),
	)
}

// 9
func parsePartByValue(str string, min, max int) (*part, error) {
	if v, err := parseValue(str, min, max); err != nil {
		return nil, err
	} else {
		return &part{value: v}, nil
	}
}

// 1,6,8
func parsePartByMultiValue(str string, min, max int) (*part, error) {
	if res, err := parseValues(strings.Split(str, ","), min, max); err != nil {
		return nil, err
	} else {
		return &part{values: res}, nil
	}
}

// 5-12
func parsePartByRange(str string, min, max int) (*part, error) {
	if res, err := parseRanges(str, min, max); err != nil {
		return nil, err
	} else {
		return &part{ranges: res}, nil
	}
}

// */2	1-6/2
func parsePartByPeriod(str string, min, max int) (*part, error) {
	t := strings.Split(str, "/")

	fraction, err := parseValue(t[1], 1, max)
	if err != nil {
		return nil, err
	}
	if t[0] == "*" {
		return &part{period: fraction}, nil
	}

	if rangeValues.MatchString(t[0]) {
		if res, err := parseRanges(t[0], min, max); err != nil {
			return nil, err
		} else {
			return &part{ranges: res, period: fraction}, nil
		}
	}

	return nil, errors.New(
		fmt.Sprintf("Invalid value format: %s", str),
	)
}

func parseValue(str string, min, max int) (*int, error) {
	if v, err := strconv.Atoi(str); err != nil {
		return nil, err
	} else if v < min || v > max {
		return nil, valueError(str, min, max)
	} else {
		return &v, nil
	}
}

func parseValues(nums []string, min, max int) (map[int]bool, error) {
	res := make(map[int]bool)

	for _, num := range nums {
		if v, err := parseValue(num, min, max); err != nil {
			return nil, err
		} else {
			res[*v] = true
		}
	}

	return res, nil
}

func parseRanges(str string, min, max int) (*[2]int, error) {
	t := strings.Split(str, "-")
	t1, err := parseValue(t[0], min, max)
	if err != nil {
		return nil, err
	}
	t2, err := parseValue(t[1], min, max)
	if err != nil {
		return nil, err
	}
	if *t1 > *t2 {
		*t1, *t2 = *t2, *t1
	}

	list := make([]string, 0, 10)
	for i := *t1; i <= *t2; i++ {
		list = append(list, strconv.Itoa(i))
	}

	return &[2]int{*t1, *t2}, nil
}

func valueError(str string, min, max int) error {
	return errors.New(
		fmt.Sprintf("The value must be greater than/equal to %d and less than/equal to %d, pass: %s", min, max, str),
	)
}
