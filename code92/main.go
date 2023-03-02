package main

import (
	"fmt"
	listnode "my_leetcode/internal/list_node"
	"reflect"
)

// 将 [curr, end) 范围的节点进行反转
func reverseLink(pre, curr *listnode.ListNode, end *listnode.ListNode) *listnode.ListNode {
	if curr == end {
		return pre
	}

	currNext := curr.Next
	curr.Next = pre

	return reverseLink(curr, currNext, end)
}

func reverseBetween(head *listnode.ListNode, left int, right int) *listnode.ListNode {
	// 反转的起始节点
	leftNode := head
	for i := 0; i < left-1; i++ {
		leftNode = leftNode.Next
	}

	// 这里找的是right指定的节点的下一个节点，这是反转的终止节点
	endNode := head
	for i := 0; i < right-1; i++ {
		endNode = endNode.Next
	}
	// 因为 right 在设定上不会越界，所以 endNode = endNode.Next 不用做空指针判断
	endNode = endNode.Next

	reverseNode := reverseLink(endNode, leftNode, endNode)

	// 拼接反转后的节点
	if left-1 <= 0 {
		// 如果开始反转的节点(left对应的节点）就是head，那么 left-1 必定小于等于 0
		// 这时候不用做拼接，直接返回reverseNode即可
		return reverseNode
	} else {
		leftPreNode := head
		// 这里找的是开始反转的节点的前一个节点
		for i := 0; i < left-1-1; i++ {
			leftPreNode = leftPreNode.Next
		}
		leftPreNode.Next = reverseNode
	}
	return head
}

func main() {
	// 反转子链的场景
	link1 := listnode.QuickCreateListNode([]int{
		1, 2, 3, 4, 5,
	})
	newLink1 := reverseBetween(link1, 2, 4)
	newList1 := newLink1.GetList()
	expectList1 := []int{1, 4, 3, 2, 5}
	if !reflect.DeepEqual(expectList1, newList1) {
		panic(fmt.Errorf("code92 error, not expect result: \nnewLink1: %+v\nexpectLink1: %+v",
			newLink1, expectList1,
		))
	}

	// 反转头尾的场景
	link2 := listnode.QuickCreateListNode([]int{
		3, 5,
	})
	newLink2 := reverseBetween(link2, 1, 2)
	newList2 := newLink2.GetList()
	expectList2 := []int{5, 3}
	if !reflect.DeepEqual(expectList2, newList2) {
		panic(fmt.Errorf("code92 error, not expect result: \nnewLink2: %+v\nexpectLink2: %+v",
			newLink2, expectList2,
		))
	}
}
