package lockedmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLockedMap_Exists(t *testing.T) {
	m := New[string, int]()
	m.Set("a", 1)
	if !m.Exists("a") {
		t.Errorf("Exists failed")
	}
}

func TestLockedMap_Set(t *testing.T) {
	m := New[string, int]()
	m.Set("a", 1)
	if v, ok := m.Get("a"); !ok || v != 1 {
		t.Errorf("Get failed")
	}
}

func TestLockedMap_Get(t *testing.T) {
	m := New[string, int]()
	m.Set("a", 1)
	if v, ok := m.Get("a"); !ok || v != 1 {
		t.Errorf("Get failed")
	}
}

func TestLockedMap_Remove(t *testing.T) {
	m := New[string, int]()
	m.Set("a", 1)
	m.Remove("a")
	if m.Exists("a") {
		t.Errorf("Remove failed")
	}
}

func TestLockedMap_RemoveAll(t *testing.T) {
	m := New[string, int]()
	m.Set("a", 1)
	m.RemoveAll()
	if m.Exists("a") {
		t.Errorf("RemoveAll failed")
	}
	assert.Equal(t, 0, m.Size())
}

func TestLockedMap_Size(t *testing.T) {
	m := New[string, int]()
	m.Set("a", 1)
	if m.Size() != 1 {
		t.Errorf("Size failed")
	}
}

func TestLockedMap_Keys(t *testing.T) {
	m := New[string, int]()
	m.Set("a", 1)
	if len(m.Keys()) != 1 {
		t.Errorf("Keys failed")
	}
}

func TestLockedMap_Values(t *testing.T) {
	m := New[string, int]()
	m.Set("a", 1)
	if len(m.Values()) != 1 {
		t.Errorf("Values failed")
	}
}

func TestLockedMap_Range(t *testing.T) {
	m := New[string, int]()
	m.Set("a", 1)
	m.Set("b", 2)
	m.Set("c", 3)
	m.Range(func(key string, value int) bool {
		return true
	})
}
