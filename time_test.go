package gokit

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

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
