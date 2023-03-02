package main

import (
	"fmt"
	listnode "my_leetcode/internal/list_node"
	"reflect"
)

var count = 0

func removeNthFromEnd(head *listnode.ListNode, n int) *listnode.ListNode {
	defer func() {
		count = 0
	}()
	return dfs(head, n)
}

func dfs(node *listnode.ListNode, n int) *listnode.ListNode {
	if node == nil {
		return node
	}
	node.Next = dfs(node.Next, n)
	count += 1
	if count == n {
		return node.Next
	}
	return node
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
