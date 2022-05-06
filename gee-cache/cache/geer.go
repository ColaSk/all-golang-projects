package cache

type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool) // 用于根据传入的key选择合适的节点
}

// http 客户端
type PeerGetter interface {
	Get(group string, key string) ([]byte, error) // 从对应的group查询缓存值
}
