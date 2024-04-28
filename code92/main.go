package main

import (
	"fmt"
	. "my_leetcode/internal/list_node"
	"reflect"
)

func reverse(pre, curr *ListNode, end *ListNode) *ListNode {
	if curr == end {
		return pre
	}
	currNext := curr.Next
	curr.Next = pre
	return reverse(curr, currNext, end)
}

// 将 [curr, end) 范围的节点进行反转
func reverseBetween(head *ListNode, left, right int) *ListNode {
	find := func(p *ListNode, index int) *ListNode {
		if index < 0 {
			return nil
		}
		for i := 0; i < index; i++ {
			if p == nil {
				return nil
			}
			p = p.Next
		}
		return p
	}

	leftPre := find(head, left-1-1)
	rightNext := find(head, right-1+1)

	if leftPre == nil {
		return reverse(rightNext, head, rightNext)
	} else {
		leftPre.Next = reverse(rightNext, leftPre.Next, rightNext)
		return head
	}
}

func reverseBetween2(head *ListNode, left int, right int) *ListNode {
	var cache []*ListNode
	for head != nil {
		cache = append(cache, head)
		head = head.Next
	}

	reverseF := func(l []*ListNode) {
		i, j := 0, len(l)-1
		for i < j {
			l[i], l[j] = l[j], l[i]
			i++
			j--
		}
	}
	reverseF(cache[left-1 : right])

	dh := ListNode{}
	dhp := &dh
	for _, n := range cache {
		n.Next = nil
		dhp.Next = n
		dhp = dhp.Next
	}
	return dh.Next
}

func main() {
	// 反转子链的场景
	link1 := QuickCreateListNode([]int{
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
	link2 := QuickCreateListNode([]int{
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
