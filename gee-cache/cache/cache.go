package cache

import (
	"cache/lru"
	"sync"
)

type cache struct {
	mu         sync.Mutex // lru 互斥锁
	lru        *lru.Cache // lru 实例
	cacheBytes int64
}

// 添加
func (c *cache) add(key string, value ByteView) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil { // 延迟初始化
		c.lru = lru.New(c.cacheBytes, nil)
	}
	c.lru.Add(key, value)
}

//获取
func (c *cache) get(key string) (value ByteView, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		return
	}

	if v, ok := c.lru.Get(key); ok {
		return v.(ByteView), ok
	}

	return
}
