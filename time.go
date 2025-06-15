package gokit

import (
	"fmt"
	"strconv"
	"strings"
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

// FormatDuration 格式化 Duration 为 Xd Yh Zm Ws，左侧的单位如果为0则不显示
func FormatDuration(d time.Duration) string {
	if d < 0 {
		return "-" + FormatDuration(-d)
	}

	days := d / (24 * time.Hour)
	d -= days * 24 * time.Hour

	hours := d / time.Hour
	d -= hours * time.Hour

	minutes := d / time.Minute
	d -= minutes * time.Minute

	seconds := d / time.Second

	parts := make([]string, 0, 4)
	if days > 0 {
		parts = append(parts, fmt.Sprintf("%dd", days))
	}
	if len(parts) > 0 || hours > 0 {
		parts = append(parts, fmt.Sprintf("%dh", hours))
	}
	if len(parts) > 0 || minutes > 0 {
		parts = append(parts, fmt.Sprintf("%dm", minutes))
	}
	parts = append(parts, fmt.Sprintf("%ds", seconds))

	return fmt.Sprint(strings.Join(parts, " "))
}
