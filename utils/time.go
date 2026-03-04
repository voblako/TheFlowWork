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

func TimeToDate(t time.Time) string {
	day := t.Day()
	month := int(t.Month())
	year := t.Year()
	dayStr := strconv.Itoa(day)
	monthStr := strconv.Itoa(month)
	yearStr := strconv.Itoa(year)
	if day < 10 {
		dayStr = "0" + dayStr
	}
	if month < 10 {
		monthStr = "0" + monthStr
	}
	return dayStr + "." + monthStr + "." + yearStr
}
