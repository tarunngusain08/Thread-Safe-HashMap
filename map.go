package main

import (
	"sync"
)

type ConcurrentMap struct {
	mu    sync.RWMutex
	items map[string]any
}

func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		items: make(map[string]any),
	}
}

func (m *ConcurrentMap) Set(key string, value any) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.items[key] = value
}

func (m *ConcurrentMap) Get(key string) (any, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	value, ok := m.items[key]
	return value, ok
}

func (m *ConcurrentMap) Delete(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.items, key)
}

func (m *ConcurrentMap) Size() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.items)
}
