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

	require.Equal(t, int64(100), Max(int64(50), int64(100)))
	require.Equal(t, 3.14, Max(3.14, 2.71))
	require.Equal(t, "z", Max("a", "z"))
}

func TestMin(t *testing.T) {
	require.Equal(t, 1, Min(1, 3))
	require.Equal(t, 1, Min(3, 1))
	require.Equal(t, 0, Min(0, 0))

	require.Equal(t, int64(50), Min(int64(50), int64(100)))
	require.Equal(t, 2.71, Min(3.14, 2.71))
	require.Equal(t, "a", Min("a", "z"))
}
