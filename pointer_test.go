package gokit

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReference(t *testing.T) {
	s := "hello"
	p := ToPtr(s) // p is a pointer to a new string, not the original string
	require.Equalf(t, &s, p, "ToPtr failed")
	require.NotSamef(t, &s, p, "ToPtr failed")

	require.Equalf(t, s, *p, "ToPtr failed")
}

func TestDereference(t *testing.T) {
	s := "hello"
	v := FromPtr(&s)
	require.Equalf(t, s, v, "FromPtr failed")
}
