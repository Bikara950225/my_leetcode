package main

import listnode "my_leetcode/internal/list_node"

func detectCycle(head *listnode.ListNode) *listnode.ListNode {
	if head == nil {
		return nil
	}
	mark := map[*listnode.ListNode]struct{}{}
	for head != nil {
		if _, ok := mark[head]; ok {
			return head
		}
		mark[head] = struct{}{}
		head = head.Next
	}
	return nil
}

func main() {

}
