package main

import (
	"fmt"
	listnode "my_leetcode/internal/list_node"
	"reflect"
)

func removeNthFromEnd(head *listnode.ListNode, n int) *listnode.ListNode {
	count := 0
	var dfs func(node *listnode.ListNode) *listnode.ListNode
	dfs = func(node *listnode.ListNode) *listnode.ListNode {
		if node == nil {
			return nil
		}
		node.Next = dfs(node.Next)
		count++
		if count == n {
			return node.Next
		}
		return node
	}
	return dfs(head)
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
