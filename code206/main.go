package main

import (
	"fmt"
	listnode "my_leetcode/internal/list_node"
	"reflect"
)

// end为闭区间
func recursive(pre, curr *listnode.ListNode, end *listnode.ListNode) *listnode.ListNode {
	if curr == end {
		return pre
	}
	currNext := curr.Next
	curr.Next = pre

	return recursive(curr, currNext, end)
}

func reverseList(head *listnode.ListNode) *listnode.ListNode {
	return recursive(nil, head, nil)
}

func main() {
	head := listnode.QuickCreateListNode([]int{1, 2, 3})

	newNode := reverseList(head)
	newList := newNode.GetList()

	expectList := []int{3, 2, 1}
	if !reflect.DeepEqual(expectList, newList) {
		panic(fmt.Errorf("not expectList: %+v", newList))
	}
}
