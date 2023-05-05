package main

import (
	"fmt"
	listnode "my_leetcode/internal/list_node"
	"reflect"
)

func deleteDuplicates(head *listnode.ListNode) *listnode.ListNode {
	return deleteDuplicatesWithRecursive(nil, head)
}

func deleteDuplicates2(head *listnode.ListNode) *listnode.ListNode {
	if head == nil {
		return nil
	}
	sp, qp := head, head.Next
	for qp != nil {
		if sp.Val != qp.Val {
			sp.Next = qp
			sp = qp
		}
		qp = qp.Next
	}
	sp.Next = nil
	return head
}

func deleteDuplicatesWithRecursive(pre, curr *listnode.ListNode) *listnode.ListNode {
	if curr == nil {
		return nil
	}
	if pre == nil {
		curr.Next = deleteDuplicatesWithRecursive(curr, curr.Next)
		return curr
	}

	if pre.Val == curr.Val {
		return deleteDuplicatesWithRecursive(pre, curr.Next)
	}
	curr.Next = deleteDuplicatesWithRecursive(curr, curr.Next)
	return curr
}

func main() {
	ln1 := listnode.QuickCreateListNode([]int{1, 1, 2})
	retLn1 := deleteDuplicates(ln1)
	if !reflect.DeepEqual(retLn1.GetList(), []int{1, 2}) {
		panic(fmt.Errorf("code83 error, retLn1 != [1, 2]: %v", retLn1))
	}

	ln2 := listnode.QuickCreateListNode([]int{})
	retLn2 := deleteDuplicates(ln2)
	var emptySlice []int
	if !reflect.DeepEqual(retLn2.GetList(), emptySlice) {
		panic(fmt.Errorf("code83 error, retLn2 != []: %v", retLn2))
	}

	ln3 := listnode.QuickCreateListNode([]int{1, 1, 2, 3, 3})
	retLn3 := deleteDuplicates2(ln3)
	if !reflect.DeepEqual(retLn3.GetList(), []int{1, 2, 3}) {
		panic(fmt.Errorf("code83 error, retLn3 != []: %v", retLn3))
	}
}
