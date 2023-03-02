package main

import listnode "my_leetcode/internal/list_node"

// 空间复杂度：O(n)
// 时间复杂度：O(n)
func isPalindrome1(head *listnode.ListNode) bool {
	h := head
	finished := false

	var recursiveCheck func(n *listnode.ListNode) bool
	recursiveCheck = func(n *listnode.ListNode) bool {
		if n == nil {
			return true
		}
		ret := recursiveCheck(n.Next)
		if !ret || finished {
			return ret
		}
		if n == h {
			finished = true
			return ret
		}

		if n.Val == h.Val {
			h = h.Next
			return true
		}
		return false
	}
	return recursiveCheck(head.Next)
}

func findMiddle(head *listnode.ListNode) *listnode.ListNode {
	sp, qp := head, head.Next
	for qp != nil {
		qp = qp.Next
		if qp != nil {
			qp = qp.Next
			sp = sp.Next
		}
	}
	return sp
}

func reverse(pre, curr, end *listnode.ListNode) *listnode.ListNode {
	if curr == end {
		return pre
	}
	currNext := curr.Next
	curr.Next = pre
	return reverse(curr, currNext, end)
}