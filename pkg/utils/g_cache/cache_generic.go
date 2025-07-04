package g_cache

import "sync"

type Cache[K comparable, V any] interface {
    Get(key K) interface{}
    Set(key K, value interface{})
    Delete(key K)
    GetValue(key K) (isExist bool, val V)
}

type mapCache[K comparable, V any] struct {
    data map[K]interface{}
    lock sync.RWMutex
}

func NewMapCache[K comparable, V any]() Cache[K, V] {
    return &mapCache[K, V]{
        data: make(map[K]interface{}),
    }
}

func (m *mapCache[K, V]) Get(key K) interface{} {
    m.lock.RLock()
    defer m.lock.RUnlock()
    return m.data[key]
}

func (m *mapCache[K, V]) GetValue(key K) (isExist bool, val V) {
    m.lock.RLock()
    defer m.lock.RUnlock()
    val, isExist = m.data[key].(V)
    return
}

func (m *mapCache[K, V]) Set(key K, value interface{}) {
    m.lock.Lock()
    defer m.lock.Unlock()
    m.data[key] = value
}

func (m *mapCache[K, V]) Delete(key K) {
    m.lock.Lock()
    defer m.lock.Unlock()
    delete(m.data, key)
}