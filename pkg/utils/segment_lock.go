package utils

import (
	"hash/fnv"
	"sync"
)

// 分段计数器结构体
type SegmentedCounter struct {
	segments []*counterSegment // 分段数组
	numSegs  int               // 分段数量
}

// 单个计数分段
type counterSegment struct {
	mu     sync.RWMutex     // 读写锁，优化读性能
	counts map[string]int64 // 存储具体键的计数
}

// 创建新的分段计数器
func NewSegmentedCounter(numSegs int) *SegmentedCounter {
	segments := make([]*counterSegment, numSegs)
	for i := 0; i < numSegs; i++ {
		segments[i] = &counterSegment{
			counts: make(map[string]int64),
		}
	}
	return &SegmentedCounter{
		segments: segments,
		numSegs:  numSegs,
	}
}

// 计算键所属的分段索引 将不同的 key 均匀映射到不同的分段（segment），这是分段锁实现中至关重要的一步
func (c *SegmentedCounter) getSegmentIndex(key string) int {
	hash := fnv.New32a()
	hash.Write([]byte(key))
	return int(hash.Sum32() % uint32(c.numSegs))
}

// 增加计数（线程安全）
func (c *SegmentedCounter) Incr(key string) {
	idx := c.getSegmentIndex(key)
	seg := c.segments[idx]

	seg.mu.Lock() // 写操作加互斥锁
	defer seg.mu.Unlock()
	seg.counts[key]++
}

// 获取计数（线程安全）
func (c *SegmentedCounter) Get(key string) int64 {
	idx := c.getSegmentIndex(key)
	seg := c.segments[idx]

	seg.mu.RLock() // 读操作加共享锁
	defer seg.mu.RUnlock()
	return seg.counts[key]
}

// 获取所有键的总计数（遍历所有分段）
func (c *SegmentedCounter) Total() map[string]int64 {
	total := make(map[string]int64)

	// 依次锁定每个分段并累加计数
	for _, seg := range c.segments {
		seg.mu.RLock()
		for key, count := range seg.counts {
			total[key] += count
		}
		seg.mu.RUnlock()
	}

	return total
}
