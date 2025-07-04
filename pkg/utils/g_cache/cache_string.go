package g_cache

import "sync"

type CacheString interface {
	Get(key string) interface{}
	Set(key string, value interface{})
	Delete(key string)
	GetValueString(key string) (isExist bool, val string)
}

type mapCacheString struct {
	data map[string]interface{}
	lock sync.RWMutex
}

func NewMapCacheString() CacheString {
	return &mapCacheString{
		data: make(map[string]interface{}),
	}
}

func (m *mapCacheString) Get(key string) interface{} {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.data[key]
}

func (m *mapCacheString) GetValueString(key string) (isExist bool, val string) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	val, isExist = m.data[key].(string)
	return
}

func (m *mapCacheString) Set(key string, value interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.data[key] = value
}

func (m *mapCacheString) Delete(key string) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.data, key)
}
