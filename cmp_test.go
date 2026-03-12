package gokit

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAbs(t *testing.T) {
	require.Equal(t, 3, Abs(3))
	require.Equal(t, 3, Abs(-3))
	require.Equal(t, 0, Abs(0))

	require.Equal(t, int64(100), Abs(int64(-100)))
	require.Equal(t, 3.14, Abs(-3.14))
	require.Equal(t, float32(2.5), Abs(float32(-2.5)))
}

func TestMax(t *testing.T) {
	require.Equal(t, 3, Max(1, 3))
	require.Equal(t, 3, Max(3, 1))
	require.Equal(t, 0, Max(0, 0))
	require.Equal(t, 42, Max(42))
	require.Equal(t, 9, Max(3, 9, 1, 8, 2))
	require.Equal(t, -1, Max(-10, -3, -7, -1))

	require.Equal(t, int64(100), Max(int64(50), int64(100)))
	require.Equal(t, int64(200), Max(int64(50), int64(100), int64(200), int64(150)))
	require.Equal(t, 3.14, Max(3.14, 2.71))
	require.Equal(t, 9.99, Max(1.23, 9.99, 4.56))
	require.Equal(t, "z", Max("a", "z"))
	require.Equal(t, "z", Max("a", "m", "z", "b"))
}

func TestMin(t *testing.T) {
	require.Equal(t, 1, Min(1, 3))
	require.Equal(t, 1, Min(3, 1))
	require.Equal(t, 0, Min(0, 0))
	require.Equal(t, 42, Min(42))
	require.Equal(t, 1, Min(3, 9, 1, 8, 2))
	require.Equal(t, -10, Min(-10, -3, -7, -1))

	require.Equal(t, int64(50), Min(int64(50), int64(100)))
	require.Equal(t, int64(50), Min(int64(200), int64(100), int64(50), int64(150)))
	require.Equal(t, 2.71, Min(3.14, 2.71))
	require.Equal(t, 1.23, Min(1.23, 9.99, 4.56))
	require.Equal(t, "a", Min("a", "z"))
	require.Equal(t, "a", Min("m", "z", "a", "b"))
}
