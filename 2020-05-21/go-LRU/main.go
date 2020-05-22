package mian

import "errors"

// Node 定义链表的节点
type Node struct {
	pre, next Node
	Key, Val  int
}

// DoubleList 定义双向列表
type DoubleList struct {
	head, tail Node
	MaxSize    int
	// Size() int
	// RemoveLast() Node
	// AddFirst()
	// Remove(x Node)
}

// Size 返回链表当前大小
func (d DoubleList) Size() int {
	return d.Size
}

// RemoveLast 移除链表末尾元素
func (d DoubleList) RemoveLast() Node {

}

// AddFirst 向链表头部增加节点
func (d DoubleList) AddFirst() {

}

// Remove 删除链表中指定节点
func (d DoubleList) Remove(n Node) {

}

// LRUCache 定义 LRU缓存
// map 是根据hash表实现的
// 链表节点有了，链表方法定义了
type LRUCache struct {
	Map  map[int]int
	Cash DoubleList
}

// Put 元素
func (lru LRUCache) Put(key, value int) {
	node := Node{
		Key: key,
		Val: value,
	}
	if lru.Map[key] != nil { // 已经有了
		// 链表删除
		lru.Cash.Remove(node)
		// 链表存为头节点
		lru.Cash.AddFirst(node)
		// map put 更新值
		lru.Map[key] = value
	} else {
		if lru.Cash.Size >= lru.Cash.MaxSize {
			last := lru.Cash.RemoveLast()
			delete(lru.Map, last.Key)
		}
		lru.Map[key] = value
		lru.Cash.AddFirst(node)
	}

}

// Get 元素
func (lru LRUCache) Get(key int) (value int, err error) {
	if lru.Map[key] == nil {
		return (-1, errors.New())
	} 
	value := lru.Map[key]
	lru.Put(key, value)  // 利用Put方法把数据提前
	return (value,nil)

}
