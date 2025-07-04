package g_cache

import "sync"

type CacheInt interface {
	Get(key int) interface{}
	Set(key int, value interface{})
	Delete(key int)
	GetValueInt(key int) (isExist bool, val int)
}

type mapCacheInt struct {
	data map[int]interface{}
	lock sync.RWMutex
}

func NewMapCacheInt() CacheInt {
	return &mapCacheInt{
		data: make(map[int]interface{}),
	}
}

func (m *mapCacheInt) Get(key int) interface{} {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.data[key]
}

func (m *mapCacheInt) GetValueInt(key int) (isExist bool, val int) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	val, isExist = m.data[key].(int)
	return
}

func (m *mapCacheInt) Set(key int, value interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.data[key] = value
}

func (m *mapCacheInt) Delete(key int) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.data, key)
}
