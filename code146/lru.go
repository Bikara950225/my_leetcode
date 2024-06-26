package main

// Least Recently Used
type LRUCache1 struct {
	cap int

	cache map[int]*DoubleLinkNode
	// recent cache
	head *DoubleLinkNode
	tail *DoubleLinkNode
}

func ConstructorLRUCache(capacity int) LRUCache1 {
	return LRUCache1{
		cap:   capacity,
		cache: make(map[int]*DoubleLinkNode, capacity),
		head:  nil,
		tail:  nil,
	}
}

func (s *LRUCache1) Get(key int) int {
	node, ok := s.cache[key]
	if !ok {
		return -1
	}
	// 刷新recent缓存
	// 将目标节点置于头部
	s.updateRecentCache(node)
	return node.val
}

func (s *LRUCache1) Put(key int, value int) {
	putNode, ok := s.cache[key]
	if ok {
		putNode.val = value
	} else {
		putNode = &DoubleLinkNode{key: key, val: value}
	}
	// 更新recent缓存
	s.updateRecentCache(putNode)
	s.cache[key] = putNode
	// 加完新节点再删除尾巴节点
	if len(s.cache) > s.cap {
		deleteNode := s.tail
		s.tail = deleteTail(s.tail)
		delete(s.cache, deleteNode.key)
	}
}

func (s *LRUCache1) updateRecentCache(newNode *DoubleLinkNode) {
	newHead, newTail := setTop(s.head, newNode)
	if newHead != nil {
		s.head = newHead
	}
	if newTail != nil {
		s.tail = newTail
	}
}
