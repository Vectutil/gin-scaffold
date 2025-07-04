package redis

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

// Set 根据Key 设置redis
func (c *redisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	err := c.client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

// Get 根据Key 获取redis
func (c *redisClient) Get(ctx context.Context, key string) (string, error) {
	val, err := c.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

// Delete 根据Key 删除redis
func (c *redisClient) Delete(ctx context.Context, key string) error {
	err := c.client.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}

// Close 关闭redis
func (c *redisClient) Close() error {
	return c.client.Close()
}

// Lock 分布式锁
func (c *redisClient) Lock(ctx context.Context, key, value string, expiration time.Duration) (bool, error) {
	return c.client.SetNX(ctx, key, value, expiration).Result()
}

// UnLock 解锁
func (c *redisClient) UnLock(ctx context.Context, key, value string) error {
	var unlockScript = redis.NewScript(`
		if redis.call("get", KEYS[1]) == ARGV[1] then
			return redis.call("del", KEYS[1])
		else
			return 0
		end
		`)

	result, err := unlockScript.Run(ctx, c.client, []string{key}, value).Int()
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("ErrLockNotHold")
	}
	return nil
}
