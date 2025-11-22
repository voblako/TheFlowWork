package utils

import (
	"strconv"
	"time"
)

// 13.04.2007
func DateToTime(date string) (time.Time, error) {

	day, err := strconv.Atoi(date[0:2])
	if err != nil {
		return time.Time{}, err
	}
	monthInt, err := strconv.Atoi(date[3:5])
	if err != nil {
		return time.Time{}, err
	}
	year, err := strconv.Atoi(date[6:10])
	if err != nil {
		return time.Time{}, err
	}
	return time.Date(year, time.Month(monthInt), day, 0, 0, 0, 0, time.Local), nil
}
