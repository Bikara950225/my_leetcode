package main

import (
	"fmt"
	. "my_leetcode/internal/list_node"
	"reflect"
)

func addTwoNumbers2(l1, l2 *ListNode) *ListNode {
	head := ListNode{}
	curr := &head

	isUp := false
	for l1 != nil || l2 != nil && isUp {
		curr.Next = &ListNode{}
		curr = curr.Next

		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		currVal := v1 + v2
		if isUp {
			currVal += 1
		}
		if currVal >= 10 {
			isUp = true
			currVal -= 10
		} else {
			isUp = false
		}
		curr.Val = currVal
	}

	return head.Next
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return bfs(l1, l2)
}

func bfs(l1, l2 *ListNode) *ListNode {
	// dummy head
	dh := ListNode{}
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

		curr.Next = &ListNode{Val: currSum}
		curr = curr.Next
	}

	return dh.Next
}

func main() {
	src1_1 := QuickCreateListNode([]int{2, 4, 3})
	src1_2 := QuickCreateListNode([]int{5, 6, 4})
	ret1 := addTwoNumbers2(src1_1, src1_2)
	ret1List := ret1.GetList()
	if !reflect.DeepEqual(ret1List, []int{7, 0, 8}) {
		panic(fmt.Errorf("code2 error, not expect result: %+v", ret1List))
	}

	src2_1 := QuickCreateListNode([]int{2, 4})
	src2_2 := QuickCreateListNode([]int{1, 6, 5})
	ret2 := addTwoNumbers2(src2_1, src2_2)
	ret2List := ret2.GetList()
	if !reflect.DeepEqual(ret2List, []int{3, 0, 6}) {
		panic(fmt.Errorf("code2 error, not expect result: %+v", ret2List))
	}
}
