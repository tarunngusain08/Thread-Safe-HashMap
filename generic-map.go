package main

import (
	"sync"
)

type GenericConcurrentMap[K comparable, V any] struct {
	mu    sync.RWMutex
	items map[K]V
}

func NewGenericConcurrentMap[K comparable, V any]() *GenericConcurrentMap[K, V] {
	return &GenericConcurrentMap[K, V]{
		items: make(map[K]V),
	}
}

func (m *GenericConcurrentMap[K, V]) Set(key K, value V) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.items[key] = value
}

func (m *GenericConcurrentMap[K, V]) Get(key K) (value V, ok bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	value, ok = m.items[key]
	return value, ok
}

func (m *GenericConcurrentMap[K, V]) Delete(key K) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.items, key)
}

func (m *GenericConcurrentMap[K, V]) Size() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.items)
}
