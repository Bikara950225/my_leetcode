package main

import (
	"fmt"
	treenode "my_leetcode/internal/tree_node"
	"reflect"
)

var (
	ret    [][]int
	buffer []int
	sum    int
)

func pathSum(root *treenode.TreeNode, targetSum int) [][]int {
	defer func() {
		ret = nil
		buffer = nil
		sum = 0
	}()

	dfs(root, targetSum)
	return ret
}

func dfs(n *treenode.TreeNode, targetSum int) {
	if n == nil {
		return
	}

	sum += n.Val
	buffer = append(buffer, n.Val)
	defer func() {
		sum -= n.Val
		buffer = buffer[:len(buffer)-1]
	}()

	if n.Left == nil && n.Right == nil && sum == targetSum {
		copyBuffer := make([]int, len(buffer))
		copy(copyBuffer, buffer)
		ret = append(ret, copyBuffer)

		return
	}

	dfs(n.Left, targetSum)
	dfs(n.Right, targetSum)
}

func main() {
	srcTree1 := &treenode.TreeNode{
		Val: 1,
		Left: &treenode.TreeNode{
			Val: 2,
		},
		Right: &treenode.TreeNode{
			Val: 3,
		},
	}
	ret1 := pathSum(srcTree1, 4)
	expectRet1 := [][]int{
		{1, 3},
	}
	if !reflect.DeepEqual(ret1, expectRet1) {
		panic(fmt.Errorf("code113 error, not expect result: %+v", expectRet1))
	}
}
