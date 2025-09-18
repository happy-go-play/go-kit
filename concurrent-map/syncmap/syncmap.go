package syncmap

import (
	"sync"
)

// Map 是一个类型安全的线程安全映射
type Map[K comparable, V any] struct {
	m sync.Map
}

func New[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{}
}

func (m *Map[K, V]) Store(key K, value V) {
	m.m.Store(key, value)
}

// Load 获取 key 对应的 value
// 如果 key 不存在，返回零值和 false
func (m *Map[K, V]) Load(key K) (V, bool) {
	val, ok := m.m.Load(key)
	if !ok {
		var zero V
		return zero, false
	}
	return val.(V), true
}

func (m *Map[K, V]) Delete(key K) {
	m.m.Delete(key)
}

// LoadOrStore 获取 key 对应的 value，如果不存在则设置为 value
// 返回实际存储的 value，以及是否已经存在
func (m *Map[K, V]) LoadOrStore(key K, value V) (V, bool) {
	actual, loaded := m.m.LoadOrStore(key, value)
	return actual.(V), loaded
}

// Range 遍历 Map 中的所有键值对
func (m *Map[K, V]) Range(f func(key K, value V) bool) {
	m.m.Range(func(k, v any) bool {
		return f(k.(K), v.(V))
	})
}
