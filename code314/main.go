package main

import (
	. "my_leetcode/internal/tree_node"
	"sort"
)

type pair struct {
	val    int
	layout int
}

func verticalOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	m := map[int][]*pair{}
	minI := 0
	var ff func(n *TreeNode, i, layout int)
	ff = func(n *TreeNode, i, layout int) {
		if n == nil {
			return
		}
		if _, ok := m[i]; !ok {
			m[i] = []*pair{}
		}
		m[i] = append(m[i], &pair{
			val:    n.Val,
			layout: layout,
		})
		minI = min(minI, i)

		ff(n.Left, i-1, layout+1)
		ff(n.Right, i+1, layout+1)
	}
	ff(root, 0, 0)

	// 遍历每个子数组，对它们进行layout排序
	for _, sub := range m {
		sort.SliceStable(sub, func(i, j int) bool {
			return sub[i].layout < sub[j].layout
		})
	}

	e := minI + len(m)
	ret := make([][]int, 0, len(m))
	for ; minI < e; minI++ {
		subRet := make([]int, len(m[minI]))
		for i, item := range m[minI] {
			subRet[i] = item.val
		}
		ret = append(ret, subRet)
	}
	return ret
}
