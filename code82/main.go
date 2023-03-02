package main

import (
	"fmt"
	listnode "my_leetcode/internal/list_node"
	"reflect"
)

func deleteDuplicates(head *listnode.ListNode) *listnode.ListNode {
	if head == nil {
		return nil
	}

	dummyHead := listnode.ListNode{}
	currNode := &dummyHead

	sp := head
	qp := head.Next
	mark := -1 << 63
	for qp != nil {
		if sp.Val != qp.Val {
			if sp.Val != mark {
				currNode.Next = &listnode.ListNode{Val: sp.Val}
				currNode = currNode.Next
			}
			sp = qp
			qp = qp.Next
			continue
		}
		mark = qp.Val
		qp = qp.Next
	}
	if mark != sp.Val {
		currNode.Next = &listnode.ListNode{Val: sp.Val}
	}
	return dummyHead.Next
}

func main() {
	src1 := listnode.QuickCreateListNode([]int{
		1, 2, 3, 3, 4, 4, 5,
	})
	retLink := deleteDuplicates(src1)
	retList := retLink.GetList()
	expectList := []int{
		1, 2, 5,
	}
	if !reflect.DeepEqual(retList, expectList) {
		panic(fmt.Errorf("code82 error, expect result: %+v", retList))
	}
}
