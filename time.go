package gokit

import (
	"fmt"
	"strconv"
	"time"
)

func ParseUnixTimestamp(timestamp string) (time.Time, error) {
	ts, err := strconv.Atoi(timestamp)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid timestamp: %v", err)
	}
	return time.Unix(int64(ts), 0), nil
}

func ParseDateTime(timeStr string) (time.Time, error) {
	// timeStr: "2021-08-31 00:00:00"
	t, err := time.Parse(time.DateTime, timeStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid time format: %v", err)
	}
	return t, nil
}

func StartOfDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

func EndOfDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 23, 59, 59, 999999999, t.Location())
}
