package syncmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStoreAndGet(t *testing.T) {
	t.Parallel()
	m := New[string, int]()

	m.Store("a", 1)
	val, ok := m.Load("a")
	require.True(t, ok)
	assert.Equal(t, 1, val)
}

func TestStoreAndGet2(t *testing.T) {
	t.Parallel()
	m := New[string, int]()

	m.Store("a", 1)
	val, ok := m.Load("a")
	require.True(t, ok)
	assert.Equal(t, 1, val)
}

func TestGetMissing(t *testing.T) {
	t.Parallel()
	m := New[string, int]()
	val, ok := m.Load("missing")
	assert.False(t, ok)
	assert.Equal(t, 0, val)
}

func TestDelete(t *testing.T) {
	t.Parallel()
	m := New[string, int]()
	m.Store("a", 1)
	m.Delete("a")
	_, ok := m.Load("a")
	assert.False(t, ok)
}

func TestLoadOrStore(t *testing.T) {
	t.Parallel()
	m := New[string, int]()

	// first time loads provided value
	val, loaded := m.LoadOrStore("a", 1)
	assert.False(t, loaded)
	assert.Equal(t, 1, val)

	// second time should not overwrite
	val, loaded = m.LoadOrStore("a", 2)
	assert.True(t, loaded)
	assert.Equal(t, 1, val)
}

func TestRange(t *testing.T) {
	t.Parallel()
	m := New[string, int]()
	m.Store("a", 1)
	m.Store("b", 2)
	m.Store("c", 3)

	sum := 0
	m.Range(func(key string, value int) bool {
		sum += value
		return true
	})
	assert.Equal(t, 6, sum)
}
