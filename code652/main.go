package main

import (
	"fmt"
	treenode "my_leetcode/internal/tree_node"
	"strings"
)

func findDuplicateSubtrees(root *treenode.TreeNode) (ret []*treenode.TreeNode) {
	// 记录出现过的子树
	pathMap := map[string]int{}

	var parseFunc func(*treenode.TreeNode) string
	parseFunc = func(n *treenode.TreeNode) string {
		if n == nil {
			return "()"
		}

		left := parseFunc(n.Left)
		right := parseFunc(n.Right)

		subTree := strings.Join([]string{
			fmt.Sprintf("(%d)", n.Val), left, right,
		}, "+")
		pathMap[subTree] += 1
		if pathMap[subTree] == 2 {
			ret = append(ret, n)
		}

		return subTree
	}
	parseFunc(root)
	return
}

func main() {
	tree1 := &treenode.TreeNode{
		Val: 1,
		Left: &treenode.TreeNode{
			Val:  2,
			Left: &treenode.TreeNode{Val: 4},
		},
		Right: &treenode.TreeNode{
			Val: 3,
			Left: &treenode.TreeNode{
				Val:  2,
				Left: &treenode.TreeNode{Val: 4},
			},
			Right: &treenode.TreeNode{Val: 4},
		},
	}
	ret1 := findDuplicateSubtrees(tree1)
	print(len(ret1))
}
