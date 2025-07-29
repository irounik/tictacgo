package db

import "sync"

type InMemoryDb[K comparable, V any] struct {
	dataMap map[K]V
	mu      sync.RWMutex
}

func NewInMemoryDb[K comparable, V any]() *InMemoryDb[K, V] {
	return &InMemoryDb[K, V]{dataMap: make(map[K]V)}
}

func (db *InMemoryDb[K, V]) Save(id K, data V) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.dataMap[id] = data
}

func (db *InMemoryDb[K, V]) Get(id K) (V, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	val, exists := db.dataMap[id]
	return val, exists
}

func (db *InMemoryDb[K, V]) Exists(id K) bool {
	db.mu.RLock()
	defer db.mu.RUnlock()
	_, exists := db.dataMap[id]
	return exists
}

func (db *InMemoryDb[K, V]) Delete(id K) {
	db.mu.Lock()
	defer db.mu.Unlock()
	delete(db.dataMap, id)
}
