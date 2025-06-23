package gokit

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestStartOfDayInLocation(t *testing.T) {
	loc, err := time.LoadLocation("Asia/Shanghai")
	require.NoError(t, err, "Failed to load location")

	{
		now := time.Date(2025, 6, 2, 1, 1, 1, 1, loc)
		startOfHour := StartOfDayInLocation(now, loc)
		expected := time.Date(2025, 6, 2, 0, 0, 0, 0, loc)
		require.Equal(t, expected, startOfHour, "Expected %s but got %s", expected, startOfHour)
	}

	{
		now := time.Date(2025, 6, 2, 1, 1, 1, 1, loc)
		t.Logf("now in %s is %s", loc, now)
		t.Logf("now in UTC is %s", now.UTC())
		startOfHour := StartOfDayInLocation(now, time.UTC)
		expected := time.Date(2025, 6, 1, 0, 0, 0, 0, time.UTC)
		require.Equal(t, expected, startOfHour, "Expected %s but got %s", expected, startOfHour)
	}
}

func TestStartOfHourInLocation(t *testing.T) {
	loc, err := time.LoadLocation("Asia/Shanghai")
	require.NoError(t, err, "Failed to load location")

	{
		now := time.Date(2025, 6, 2, 1, 1, 1, 1, loc)
		startOfHour := StartOfHourInLocation(now, loc)
		expected := time.Date(2025, 6, 2, 1, 0, 0, 0, loc)
		require.Equal(t, expected, startOfHour, "Expected %s but got %s", expected, startOfHour)
	}

	{
		now := time.Date(2025, 6, 2, 1, 1, 1, 1, loc)
		t.Logf("now in %s is %s", loc, now)
		t.Logf("now in UTC is %s", now.UTC())
		startOfHour := StartOfHourInLocation(now, time.UTC)
		expected := time.Date(2025, 6, 1, 17, 0, 0, 0, time.UTC)
		require.Equal(t, expected, startOfHour, "Expected %s but got %s", expected, startOfHour)
	}
}

func TestEndOfHourInLocation(t *testing.T) {
	loc, err := time.LoadLocation("Asia/Shanghai")
	require.NoError(t, err, "Failed to load location")

	{
		now := time.Date(2025, 6, 2, 1, 1, 1, 1, loc)
		endOfHour := EndOfHourInLocation(now, loc)
		expected := time.Date(2025, 6, 2, 1, 59, 59, 999999999, loc)
		require.Equal(t, expected, endOfHour, "Expected %s but got %s", expected, endOfHour)
	}

	{
		now := time.Date(2025, 6, 2, 1, 1, 1, 1, loc)
		t.Logf("now in %s is %s", loc, now)
		t.Logf("now in UTC is %s", now.UTC())
		endOfHour := EndOfHourInLocation(now, time.UTC)
		expected := time.Date(2025, 6, 1, 17, 59, 59, 999999999, time.UTC)
		require.Equal(t, expected, endOfHour, "Expected %s but got %s", expected, endOfHour)
	}
}

func TestFormatDuration(t *testing.T) {
	{
		d := 10 * 24 * time.Hour
		result := FormatDuration(d)
		expected := "10d 0h 0m 0s"
		require.Equal(t, expected, result, "Expected %s but got %s", expected, result)
	}

	{
		d := 2 * time.Hour
		result := FormatDuration(d)
		expected := "2h 0m 0s"
		require.Equal(t, expected, result, "Expected %s but got %s", expected, result)
	}

	{
		d := 30 * time.Minute
		result := FormatDuration(d)
		expected := "30m 0s"
		require.Equal(t, expected, result, "Expected %s but got %s", expected, result)
	}

	{
		d := 45 * time.Second
		result := FormatDuration(d)
		expected := "45s"
		require.Equal(t, expected, result, "Expected %s but got %s", expected, result)
	}

	{
		d := 0 * time.Second
		result := FormatDuration(d)
		expected := "0s"
		require.Equal(t, expected, result, "Expected %s but got %s", expected, result)
	}

	{
		d := 1*time.Hour + 30*time.Minute + 45*time.Second
		result := FormatDuration(d)
		expected := "1h 30m 45s"
		require.Equal(t, expected, result, "Expected %s but got %s", expected, result)
	}

	{
		d := 11*24*time.Hour + 1*time.Hour + 30*time.Minute + 45*time.Second
		result := FormatDuration(d)
		expected := "11d 1h 30m 45s"
		require.Equal(t, expected, result, "Expected %s but got %s", expected, result)
	}

	{
		d := 11*24*time.Hour + 0*time.Hour + 30*time.Minute + 45*time.Second
		result := FormatDuration(d)
		expected := "11d 0h 30m 45s"
		require.Equal(t, expected, result, "Expected %s but got %s", expected, result)
	}

}
