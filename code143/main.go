package main

import (
	"fmt"
	listnode "my_leetcode/internal/list_node"
	"reflect"
)

func reorderList(head *listnode.ListNode) {
	var nodeCache []*listnode.ListNode
	node := head
	for node != nil {
		nodeCache = append(nodeCache, node)
		// 加入到cache中，链表的原顺序通过cache的list维护，这里删掉它们的Next链路
		nodeNext := node.Next
		node.Next = nil

		node = nodeNext
	}

	dummyHead := listnode.ListNode{}
	node = &dummyHead
	for head := 0; head < len(nodeCache); head++ {
		tail := len(nodeCache) - head - 1
		if head >= tail {
			if head == tail {
				// 链表节点数为奇数时，最后一个节点就是 head == tail
				node.Next = nodeCache[head]
			}
			break
		}
		nodeCache[head].Next = nodeCache[tail]
		node.Next = nodeCache[head]
		node = nodeCache[tail]
	}
	// 替换head指针的内容，不要替换指针
	*head = *dummyHead.Next
}

func main() {
	src1 := listnode.QuickCreateListNode([]int{1, 2, 3, 4})
	reorderList(src1)
	ret1 := src1.GetList()
	expectRet1 := []int{1, 4, 2, 3}
	if !reflect.DeepEqual(ret1, expectRet1) {
		panic(fmt.Errorf("code143 error, not expect result: %+v", ret1))
	}

	src2 := listnode.QuickCreateListNode([]int{1, 2, 3, 4, 5})
	reorderList(src2)
	ret2 := src2.GetList()
	expectRet2 := []int{1, 5, 2, 4, 3}
	if !reflect.DeepEqual(ret2, expectRet2) {
		panic(fmt.Errorf("code143 error, not expect result: %+v", ret2))
	}
}
