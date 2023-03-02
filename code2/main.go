package main

import (
	"fmt"
	listnode "my_leetcode/internal/list_node"
	"reflect"
)

func addTwoNumbers(l1 *listnode.ListNode, l2 *listnode.ListNode) *listnode.ListNode {
	return bfs(l1, l2)
}

func bfs(l1, l2 *listnode.ListNode) *listnode.ListNode {
	// dummy head
	dh := listnode.ListNode{}
	curr := &dh

	isUp := false
	for l1 != nil || l2 != nil || isUp {
		v1 := 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		v2 := 0
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		currSum := v1 + v2
		if isUp {
			currSum++
		}
		isUp = false
		if currSum >= 10 {
			currSum -= 10
			isUp = true
		}

		curr.Next = &listnode.ListNode{Val: currSum}
		curr = curr.Next
	}

	return dh.Next
}

func main() {
	src1_1 := listnode.QuickCreateListNode([]int{2, 4, 3})
	src1_2 := listnode.QuickCreateListNode([]int{5, 6, 4})
	ret1 := addTwoNumbers(src1_1, src1_2)
	ret1List := ret1.GetList()
	if !reflect.DeepEqual(ret1List, []int{7, 0, 8}) {
		panic(fmt.Errorf("code2 error, not expect result: %+v", ret1List))
	}

	src2_1 := listnode.QuickCreateListNode([]int{2, 4})
	src2_2 := listnode.QuickCreateListNode([]int{1, 6, 5})
	ret2 := addTwoNumbers(src2_1, src2_2)
	ret2List := ret2.GetList()
	if !reflect.DeepEqual(ret2List, []int{3, 0, 6}) {
		panic(fmt.Errorf("code2 error, not expect result: %+v", ret2List))
	}
}
