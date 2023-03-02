package main

type DoubleLinkNode struct {
	key int
	val int

	pre  *DoubleLinkNode
	next *DoubleLinkNode
}

func setTop(head, target *DoubleLinkNode) (nh *DoubleLinkNode, nt *DoubleLinkNode) {
	// 头节点为nil
	if head == nil {
		target.pre, target.next = nil, nil
		nh, nt = target, target
		return
	}
	// target即为head节点，不处理
	if head == target {
		return
	}
	// 将target从链表中删除, 将前一个节点和下一个节点做拼接
	if target.pre != nil {
		target.pre.next = target.next
	}
	if target.next != nil {
		target.next.pre	= target.pre
	} else {
		// target为尾部节点
		nt = target.pre
	}

	target.next = head
	head.pre = target
	nh = target

	return
}

func deleteTail(tail *DoubleLinkNode) (nt *DoubleLinkNode) {
	if tail == nil || tail.next != nil {
		return
	}
	nt = tail.pre
	if nt != nil {
		nt.next = nil
	}
	return
}

func (s *DoubleLinkNode) getListFromNext() []int {
	var ret []int
	n := s
	for n != nil {
		ret = append(ret, n.val)
		n = n.next
	}
	return ret
}

func (s *DoubleLinkNode) getListFromPre() []int {
	var ret []int
	n := s
	for n != nil {
		ret = append(ret, n.val)
		n = n.pre
	}
	return ret
}
