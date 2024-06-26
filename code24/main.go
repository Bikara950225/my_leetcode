package main

import (
	"fmt"
	listnode "my_leetcode/internal/list_node"
	"reflect"
)

func swapPairs(head *listnode.ListNode) *listnode.ListNode {
	var count int

	var dfsFunc func(pre, curr *listnode.ListNode) *listnode.ListNode
	dfsFunc = func(pre, curr *listnode.ListNode) *listnode.ListNode {
		if curr == nil {
			return pre
		}
		count++

		if count&1 == 1 {
			return dfsFunc(curr, curr.Next)
		} else {
			currNext := curr.Next
			curr.Next = pre
			pre.Next = dfsFunc(nil, currNext)
			return curr
		}
	}

	return dfsFunc(nil, head)
}

func swapPairs2(head *listnode.ListNode) *listnode.ListNode {
	count := 0
	var dfs func(pre, curr *listnode.ListNode) *listnode.ListNode
	dfs = func(pre, curr *listnode.ListNode) *listnode.ListNode {
		if curr == nil {
			return nil
		}

		count++
		if count&1 == 1 {
			return dfs(curr, curr.Next)
		} else {
			currNext := curr.Next
			curr.Next = pre
			pre.Next = dfs(nil, currNext)
			return curr
		}
	}
	return dfs(nil, head)
}

func main() {
	ret1 := swapPairs2(listnode.QuickCreateListNode([]int{1, 2, 3, 4}))
	ret1List := ret1.GetList()
	if !reflect.DeepEqual(ret1List, []int{2, 1, 4, 3}) {
		panic(fmt.Errorf("code24 error, ret1 not expect: %v", ret1List))
	}
}
