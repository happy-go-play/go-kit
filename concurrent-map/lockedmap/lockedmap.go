package lockedmap

import (
	"sync"

	"golang.org/x/exp/maps"
)

// LockedMap Thread-safe map with generics support
type LockedMap[K comparable, V any] struct {
	data map[K]V
	mu   sync.RWMutex
}

func New[K comparable, V any]() *LockedMap[K, V] {
	return &LockedMap[K, V]{
		data: make(map[K]V),
		mu:   sync.RWMutex{},
	}
}

func (m *LockedMap[K, V]) Exists(key K) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()

	_, ok := m.data[key]
	return ok
}

func (m *LockedMap[K, V]) Get(key K) (V, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	value, ok := m.data[key]
	if !ok {
		return value, false
	}

	return value, true
}

func (m *LockedMap[K, V]) Set(key K, value V) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.data[key] = value
}

func (m *LockedMap[K, V]) Remove(key K) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.data, key)
}

func (m *LockedMap[K, V]) RemoveAll() {
	m.mu.Lock()
	defer m.mu.Unlock()

	//clear(m.data) // since go 1.21
	m.data = make(map[K]V)
}

func (m *LockedMap[K, V]) Size() int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return len(m.data)
}

func (m *LockedMap[K, V]) Keys() []K {
	m.mu.Lock()
	defer m.mu.Unlock()

	return maps.Keys(m.data)
}

func (m *LockedMap[K, V]) Values() []V {
	m.mu.Lock()
	defer m.mu.Unlock()

	return maps.Values(m.data)
}

func (m *LockedMap[K, V]) Range(f func(key K, value V) bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for key, value := range m.data {
		if !f(key, value) {
			break
		}
	}
}
