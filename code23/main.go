package code23

import . "my_leetcode/internal/list_node"

func mergeKLists(lists []*ListNode) *ListNode {
	// N路归并
	dh := ListNode{}
	dhp := &dh

	for {
		// 用一个cache记录当次循环的非nil节点，比对它们最小值，最小的加入到dhp中，并且将它的指针指向Next
		cache := []**ListNode{}

		for i := range lists {
			if lists[i] == nil {
				continue
			}
			cache = append(cache, &lists[i])
		}
		// lists中所有链表被遍历到nil时，break
		if len(cache) <= 0 {
			break
		}

		var minNode **ListNode
		for _, n := range cache {
			if minNode == nil {
				minNode = n
			} else {
				if (**minNode).Val > (**n).Val {
					minNode = n
				}
			}
		}
		dhp.Next = *minNode
		dhp = dhp.Next
		*minNode = (*minNode).Next
	}
	return dh.Next
}
