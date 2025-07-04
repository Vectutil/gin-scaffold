package redis

import (
	"context"
	"fmt"
	"gin-scaffold/internal/config"
	"gin-scaffold/pkg/logger"
	"sync"
	"testing"
	"time"
)

func TestConnect(t *testing.T) {
	// 初始化配置和日志
	config.InitConfig("/home/breelk/workspace/src/gin-scaffold/config.yaml")
	logger.InitLogger()
	Init()
	GetClient().Set(context.Background(), "key1", "test1", 30*time.Second)
	defer func() {

	}()
}

func TestRedisLock(t *testing.T) {
	// 初始化配置和日志
	config.InitConfig("/home/breelk/workspace/src/gin-scaffold/config.yaml")
	logger.InitLogger()
	Init()
	rdbClient := GetClient()

	wg := sync.WaitGroup{}

	shareValue := 1
	ctx := context.Background()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
				logger.Logger.Info("任务结束")
				rdbClient.UnLock(ctx, "lock1", "lock1")
			}()
			// 自旋锁
			for {
				ok, err := rdbClient.Lock(ctx, "lock1", "lock1", 30*time.Second)
				if err != nil {
					time.Sleep(3 * time.Second)
					logger.Logger.Error("redis 数据库获取数据失败")
					continue
				}
				if !ok {

					time.Sleep(3 * time.Second)
					logger.Logger.Error("redis 获取锁失败")
					continue
				}
				// 获取锁
				break
			}

			logger.Logger.Info(fmt.Sprintf("src value %d ,after value %d", shareValue, shareValue+1))
			shareValue++
		}()
		wg.Wait()
	}
	//  2025-07-04T23:57:08.920+0800	INFO	redis/eample_test.go:62	src value 1 ,after value 2
	//  2025-07-04T23:57:08.920+0800	INFO	redis/eample_test.go:41	任务结束
	//  2025-07-04T23:57:11.922+0800	ERROR	redis/eample_test.go:55	redis 获取锁失败
	//  2025-07-04T23:57:11.922+0800	INFO	redis/eample_test.go:62	src value 2 ,after value 3
	//  2025-07-04T23:57:11.922+0800	INFO	redis/eample_test.go:41	任务结束
	//  2025-07-04T23:57:11.922+0800	INFO	redis/eample_test.go:62	src value 3 ,after value 4
	//  2025-07-04T23:57:11.923+0800	INFO	redis/eample_test.go:41	任务结束
	//  2025-07-04T23:57:11.923+0800	INFO	redis/eample_test.go:62	src value 4 ,after value 5
	//  2025-07-04T23:57:11.923+0800	INFO	redis/eample_test.go:41	任务结束
	//  2025-07-04T23:57:14.925+0800	ERROR	redis/eample_test.go:55	redis 获取锁失败
	//  2025-07-04T23:57:14.925+0800	INFO	redis/eample_test.go:62	src value 5 ,after value 6
}

func TestChanel(t *testing.T) {
	// 初始化配置和日志
	config.InitConfig("/home/breelk/workspace/src/gin-scaffold/config.yaml")
	logger.InitLogger()
	var wg sync.WaitGroup
	shareValue := 1
	lockChan := make(chan struct{}, 1) // 在项目中是全局的chan

	lockChan <- struct{}{} // 获取锁-初始化
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
				logger.Logger.Info("任务结束")
				// 释放锁
				lockChan <- struct{}{}
			}()
			// 获取锁
			<-lockChan

			logger.Logger.Info(fmt.Sprintf("src value %d, after value %d", shareValue, shareValue+1))
			shareValue++
		}()
	}
	wg.Wait()

	//  2025-07-04T23:56:38.307+0800	INFO	redis/eample_test.go:102	src value 1, after value 2
	//  2025-07-04T23:56:38.308+0800	INFO	redis/eample_test.go:95	任务结束
	//  2025-07-04T23:56:38.308+0800	INFO	redis/eample_test.go:102	src value 2, after value 3
	//  2025-07-04T23:56:38.308+0800	INFO	redis/eample_test.go:95	任务结束
	//  2025-07-04T23:56:38.308+0800	INFO	redis/eample_test.go:102	src value 3, after value 4
	//  2025-07-04T23:56:38.308+0800	INFO	redis/eample_test.go:95	任务结束
	//  2025-07-04T23:56:38.308+0800	INFO	redis/eample_test.go:102	src value 4, after value 5
	//  2025-07-04T23:56:38.308+0800	INFO	redis/eample_test.go:95	任务结束
	//  2025-07-04T23:56:38.308+0800	INFO	redis/eample_test.go:102	src value 5, after value 6

}
