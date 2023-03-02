package main

import (
	"fmt"
	listnode "my_leetcode/internal/list_node"
	"reflect"
)

func mergeTwoLists(list1 *listnode.ListNode, list2 *listnode.ListNode) *listnode.ListNode {
	dummyHead := listnode.ListNode{}
	curr := &dummyHead

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}

	for list1 != nil {
		curr.Next = list1
		curr = curr.Next
		list1 = list1.Next
	}
	for list2 != nil {
		curr.Next = list2
		curr = curr.Next
		list2 = list2.Next
	}
	return dummyHead.Next
}

func main() {
	l1 := listnode.QuickCreateListNode([]int{1, 2, 4})
	l2 := listnode.QuickCreateListNode([]int{1, 3, 4})

	newLink := mergeTwoLists(l1, l2)
	newList := newLink.GetList()
	expectList := []int{1, 1, 2, 3, 4, 4}
	if !reflect.DeepEqual(newList, expectList) {
		panic(fmt.Errorf("not expectList: %+v", newList))
	}
}
