package list_node

type ListNode struct {
	Val  int
	Next *ListNode
}

func QuickCreateListNode(src []int) *ListNode {
	dummyHead := ListNode{}
	curr := &dummyHead
	for _, item := range src {
		curr.Next = &ListNode{Val: item}
		curr = curr.Next
	}
	return dummyHead.Next
}

func (s *ListNode) GetList() []int {
	curr := s
	var ret []int
	for curr != nil {
		ret = append(ret, curr.Val)
		curr = curr.Next
	}
	return ret
}
