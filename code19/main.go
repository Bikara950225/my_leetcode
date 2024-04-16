package main

import (
	"fmt"
	listnode "my_leetcode/internal/list_node"
	"reflect"
)

func removeNthFromEnd(head *listnode.ListNode, n int) *listnode.ListNode {
	count := 0
	var dfs func(node *listnode.ListNode, n int) *listnode.ListNode
	dfs = func(node *listnode.ListNode, n int) *listnode.ListNode {
		if node == nil {
			return nil
		}

		next := dfs(node.Next, n)
		count++
		if count == n {
			return next
		} else {
			node.Next = next
			return node
		}
	}
	return dfs(head, n)
}

func main() {
	srcLink := listnode.QuickCreateListNode([]int{
		1, 2, 3, 4, 5,
	})

	ret := removeNthFromEnd(srcLink, 2)
	retList := ret.GetList()
	expectRetList := []int{
		1, 2, 3, 5,
	}
	if !reflect.DeepEqual(retList, expectRetList) {
		panic(fmt.Errorf("code19 error, not expect result: %+v", ret))
	}
}
