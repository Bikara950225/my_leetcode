package main

type entry struct {
	k, v      int
	pre, next *entry
}

type LRUCache struct {
	cap, len   int
	head, tail *entry
	cache      map[int]*entry
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:   capacity,
		cache: make(map[int]*entry),
	}
}

func (s *LRUCache) Get(key int) int {
	e, ok := s.cache[key]
	if !ok {
		return -1
	}
	// 移动到头
	s.removeHead(e)
	return e.v
}

func (s *LRUCache) Put(key, value int) {
	e, ok := s.cache[key]
	if ok {
		e.v = value
		s.removeHead(e)
		return
	}
	// 淘汰
	s.removeTail()

	e = &entry{k: key, v: value}

	head := s.head
	s.head = e
	e.next = head
	if head != nil {
		head.pre = e
	}
	s.len++
	s.cache[e.k] = e
	if s.len == 1 {
		s.tail = e
	}
}

func (s *LRUCache) removeTail() {
	if s.len < s.cap {
		return
	}

	tail := s.tail
	isHead := s.head == tail
	tailPre := tail.pre
	// 改变tail指针指向
	s.tail = tailPre
	if tailPre != nil {
		tailPre.next = nil
	}
	// 如果删除的是头，头置空
	if isHead {
		s.head = nil
	}
	s.len--
	delete(s.cache, tail.k)
}

func (s *LRUCache) removeHead(e *entry) {
	if s.head == e {
		return
	}
	isTail := s.tail == e
	epre := e.pre
	enext := e.next
	head := s.head

	// 截断
	epre.next = enext
	if enext != nil {
		enext.pre = epre
	}
	// 移动到头
	s.head = e
	e.pre = nil
	e.next = head
	if head != nil {
		head.pre = e
	}
	if isTail {
		s.tail = epre
	}
}

func main() {

}
