package lru

import "container/list"

type Cache struct {
	maxBytes  int64 //最大内存
	nbytes    int64 // 已使用内存
	ll        *list.List
	cache     map[string]*list.Element
	OnEvicted func(key string, value Value)
}

// 双向链表节点的数据类型
type entry struct {
	key   string //字典键
	value Value  //实现 value 接口的任意数据类型
}

// Value 接口
type Value interface {
	Len() int
}

func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

// 查找功能
func (c *Cache) Get(key string) (Value, bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return nil, false
}

// 缓存淘汰
func (c *Cache) RemoveOldest() {
	ele := c.ll.Back() //取到队首节点并删除
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)                               // 存储的是任意类型需要转换类型
		delete(c.cache, kv.key)                                //删除映射关系
		c.nbytes -= int64(len(kv.key)) + int64(kv.value.Len()) // 更新内存
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value) // 调用回调
		}
	}
}

// 新增或更新
func (c *Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.nbytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		ele := c.ll.PushFront(&entry{key, value})
		c.cache[key] = ele
		c.nbytes += int64(len(key)) + int64(value.Len())
	}
	// 执行缓存淘汰
	for c.maxBytes != 0 && c.maxBytes < c.nbytes {
		c.RemoveOldest()
	}
}

func (c *Cache) Len() int {
	return c.ll.Len()
}
