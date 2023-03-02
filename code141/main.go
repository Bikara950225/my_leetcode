package main

import (
	"fmt"
	listnode "my_leetcode/internal/list_node"
)

func hasCycle(head *listnode.ListNode) bool {
	if head == nil {
		return false
	}

	sPointer := head
	qPointer := head
	for qPointer != nil {
		qPointer = qPointer.Next
		if sPointer == qPointer {
			return true
		}

		if qPointer == nil {
			return false
		}
		sPointer = sPointer.Next
		qPointer = qPointer.Next
		if sPointer == qPointer {
			return true
		}
	}
	return false
}

func main() {
	srcLink := listnode.QuickCreateListNode([]int{1,2,3,4,5})

	isCycle := hasCycle(srcLink)
	if isCycle != false {
		panic(fmt.Errorf("code141 error, isn't expect result: %v", isCycle))
	}
}
