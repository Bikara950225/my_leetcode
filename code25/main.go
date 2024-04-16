package main

import (
	"fmt"
	listnode "my_leetcode/internal/list_node"
	"reflect"
)

// 反转链表
func rev(pre, curr *listnode.ListNode, end *listnode.ListNode) *listnode.ListNode {
	if curr == end {
		return pre
	}
	currNext := curr.Next
	curr.Next = pre

	return rev(curr, currNext, end)
}

func reverseKGroup(head *listnode.ListNode, k int) *listnode.ListNode {
	var dfs func(mark, curr *listnode.ListNode, count, k int) *listnode.ListNode
	dfs = func(mark, curr *listnode.ListNode, count, k int) *listnode.ListNode {
		if curr == nil {
			return mark
		}

		count++
		if count == 1 {
			mark = curr
		}
		if count == k {
			endNode := curr.Next
			newNode := rev(nil, mark, endNode)
			mark.Next = dfs(nil, endNode, 0, k)
			return newNode
		} else {
			return dfs(curr, curr.Next, count, k)
		}
	}
	return dfs(nil, head, 0, k)
}

func main() {
	link := listnode.QuickCreateListNode([]int{1, 2, 3, 4, 5})

	newLink := reverseKGroup(link, 2)
	newList := newLink.GetList()
	expectList := []int{2, 1, 4, 3, 5}
	if !reflect.DeepEqual(newList, expectList) {
		panic(fmt.Errorf("reverseKGroup() error, newList != expectList: %v != %v",
			newList, expectList,
		))
	}
}
