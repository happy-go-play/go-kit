package gokit

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var (
	UTC1   = time.FixedZone("UTC+1", 1*3600)    // UTC+01
	UTC2   = time.FixedZone("UTC+2", 2*3600)    // UTC+02
	UTC3   = time.FixedZone("UTC+3", 3*3600)    // UTC+03
	UTC4   = time.FixedZone("UTC+4", 4*3600)    // UTC+04
	UTC5   = time.FixedZone("UTC+5", 5*3600)    // UTC+05
	UTC6   = time.FixedZone("UTC+6", 6*3600)    // UTC+06
	UTC7   = time.FixedZone("UTC+7", 7*3600)    // UTC+07
	UTC8   = time.FixedZone("UTC+8", 8*3600)    // UTC+08
	UTC9   = time.FixedZone("UTC+9", 9*3600)    // UTC+09
	UTC10  = time.FixedZone("UTC+10", 10*3600)  // UTC+10
	UTC11  = time.FixedZone("UTC+11", 11*3600)  // UTC+11
	UTC12  = time.FixedZone("UTC+12", 12*3600)  // UTC+12
	UTCm1  = time.FixedZone("UTC-1", -1*3600)   // UTC-01 UTC minus 1 hour
	UTCm2  = time.FixedZone("UTC-2", -2*3600)   // UTC-02 UTC minus 2 hours
	UTCm3  = time.FixedZone("UTC-3", -3*3600)   // UTC-03 UTC minus 3 hours
	UTCm4  = time.FixedZone("UTC-4", -4*3600)   // UTC-04 UTC minus 4 hours
	UTCm5  = time.FixedZone("UTC-5", -5*3600)   // UTC-05 UTC minus 5 hours
	UTCm6  = time.FixedZone("UTC-6", -6*3600)   // UTC-06 UTC minus 6 hours
	UTCm7  = time.FixedZone("UTC-7", -7*3600)   // UTC-07 UTC minus 7 hours
	UTCm8  = time.FixedZone("UTC-8", -8*3600)   // UTC-08 UTC minus 8 hours
	UTCm9  = time.FixedZone("UTC-9", -9*3600)   // UTC-09 UTC minus 9 hours
	UTCm10 = time.FixedZone("UTC-10", -10*3600) // UTC-10 UTC minus 10 hours
	UTCm11 = time.FixedZone("UTC-11", -11*3600) // UTC-11 UTC minus 11 hours
	UTCm12 = time.FixedZone("UTC-12", -12*3600) // UTC-12 UTC minus 12 hours
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

func StartOfDayInLocation(t time.Time, loc *time.Location) time.Time {
	y, m, d := t.In(loc).Date()
	return time.Date(y, m, d, 0, 0, 0, 0, loc)
}

func EndOfDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 23, 59, 59, 999999999, t.Location())
}

func EndOfDayInLocation(t time.Time, loc *time.Location) time.Time {
	y, m, d := t.In(loc).Date()
	return time.Date(y, m, d, 23, 59, 59, 999999999, loc)
}

func StartOfHour(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), 0, 0, 0, t.Location())
}

func StartOfHourInLocation(t time.Time, loc *time.Location) time.Time {
	timeInLoc := t.In(loc)
	y, m, d := timeInLoc.Date()
	return time.Date(y, m, d, timeInLoc.Hour(), 0, 0, 0, loc)
}

func EndOfHour(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), 59, 59, 999999999, t.Location())
}

func EndOfHourInLocation(t time.Time, loc *time.Location) time.Time {
	timeInLoc := t.In(loc)
	y, m, d := timeInLoc.Date()
	return time.Date(y, m, d, timeInLoc.Hour(), 59, 59, 999999999, loc)
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
