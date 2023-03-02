package main

import (
	"fmt"
	listnode "my_leetcode/internal/list_node"
	"reflect"
)

// 反转
func rev(pre, curr *listnode.ListNode, end *listnode.ListNode) *listnode.ListNode{
	if curr == end {
		return pre
	}
	currNext := curr.Next
	curr.Next = pre

	return rev(curr, currNext, end)
}

func dfs(mark, curr *listnode.ListNode, count, k int) *listnode.ListNode {
	if curr == nil {
		return mark
	}

	count += 1
	if count == 1 {
		mark = curr
	}

	if count == k {
		// 记下curr.Next, 作为反转时的终止指针
		endNode := curr.Next
		revNewNode := rev(nil, mark, endNode)
		// 经过反转后，mark节点就在末端了，它的next需要递归地进行拼接
		mark.Next = dfs(nil, endNode, 0, k)

		return revNewNode
	}

	return dfs(mark, curr.Next, count, k)
}

func reverseKGroup(head *listnode.ListNode, k int) *listnode.ListNode {
	return dfs(nil, head, 0, k)
}

func main() {
	link := listnode.QuickCreateListNode([]int{1,2,3,4,5})

	newLink := reverseKGroup(link, 2)
	newList := newLink.GetList()
	expectList := []int{2,1,4,3,5}
	if !reflect.DeepEqual(newList, expectList) {
		panic(fmt.Errorf("reverseKGroup() error, newList != expectList: %v != %v",
			newList, expectList,
		))
	}
}
