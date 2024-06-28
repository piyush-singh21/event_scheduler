package auth

import (
	"time"
)

func ParseDate(StringTime string) (time.Time, error) {
	layout := "2006-01-02T15:04:05"
	parseTime, err := time.Parse(layout, StringTime)
	if err != nil {
		return time.Time{}, err
	}
	return parseTime, nil
}
