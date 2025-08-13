package utils

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestSegmentTest(t *testing.T) {
	// 创建一个包含16个分段的计数器（通常与CPU核心数匹配）
	counter := NewSegmentedCounter(16)

	// 模拟100个并发goroutine同时访问不同的URL
	const numGoroutines = 100
	const numOperations = 1000 // 每个goroutine执行1000次操作
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// 预定义一些URL用于测试
	urls := []string{
		"https://example.com/home",
		"https://example.com/about",
		"https://example.com/blog",
		"https://example.com/contact",
		"https://example.com/products",
	}

	// 启动并发测试
	startTime := time.Now()
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			// 每个goroutine随机选择URL进行计数
			for j := 0; j < numOperations; j++ {
				idx := rand.Intn(len(urls))
				counter.Incr(urls[idx])

				// 随机进行一些查询操作
				if j%10 == 0 {
					counter.Get(urls[idx])
				}
			}
		}()
	}

	// 等待所有goroutine完成
	wg.Wait()
	duration := time.Since(startTime)

	// 输出结果
	fmt.Printf("所有操作完成，耗时: %v\n", duration)
	fmt.Println("各URL访问次数:")
	total := counter.Total()
	for url, count := range total {
		fmt.Printf("  %s: %d\n", url, count)
	}
}
