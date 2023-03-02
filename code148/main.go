package main

import (
	"fmt"
	listnode "my_leetcode/internal/list_node"
	"reflect"
)

func sortList(head *listnode.ListNode) *listnode.ListNode {
	return dfs(head, nil)
}

func dfs(head, tail *listnode.ListNode) *listnode.ListNode {
	if head == nil {
		return nil
	}
	if head.Next == tail {
		head.Next = nil
		return head
	}

	// 链表中的中间节点
	middle := findMiddle(head, tail)
	// dfs处理开闭区间[head, middle)，所以不用取middle.Next
	return merge(dfs(head, middle), dfs(middle, tail))
}

func merge(n1, n2 *listnode.ListNode) *listnode.ListNode {
	// dummy head
	dh := listnode.ListNode{}
	node := &dh

	for n1 != nil && n2 != nil {
		if n1.Val < n2.Val {
			// 这里不用复制处理
			node.Next = n1
			n1 = n1.Next
		} else {
			// 这里不用复制处理
			node.Next = n2
			n2 = n2.Next
		}
		node = node.Next
	}
	for n1 != nil {
		node.Next = n1
		n1 = n1.Next
		node = node.Next
	}
	for n2 != nil {
		node.Next = n2
		n2 = n2.Next
		node = node.Next
	}
	return dh.Next
}

func findMiddle(head, tail *listnode.ListNode) *listnode.ListNode {
	sp := head
	qp := head
	for qp != tail {
		qp = qp.Next
		if qp == tail {
			break
		}
		qp = qp.Next
		sp = sp.Next
	}
	return sp
}

func main() {
	link1 := listnode.QuickCreateListNode([]int{5, 4, 3, 2, 1})
	retLink1 := sortList(link1)
	retList1 := retLink1.GetList()
	expectList1 := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(retList1, expectList1) {
		panic(fmt.Errorf("code148 error: not expect result: %+v", retList1))
	}
}
