package main

import (
	"fmt"
	treenode "my_leetcode/internal/tree_node"
	"reflect"
)

func LevelOrderByDFS(root *treenode.TreeNode) [][]int {
	var ret [][]int
	dfs(root, &ret, 0)

	return ret
}

func dfs(node *treenode.TreeNode, ret *[][]int, level int) {
	if node == nil {
		return
	}

	currVal := node.Val
	if level >= len(*ret) {
		*ret = append(*ret, []int{})
	}
	(*ret)[level] = append((*ret)[level], currVal)

	dfs(node.Left, ret, level+1)
	dfs(node.Right, ret, level+1)
}

func LevelOrderByBFS(root *treenode.TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var ret [][]int
	nodeQueue := []*treenode.TreeNode{root}
	for len(nodeQueue) > 0 {
		layoutSize := len(nodeQueue)
		layoutList := make([]int, 0, layoutSize)
		for i := 0; i < layoutSize; i++ {
			// 出队
			n := nodeQueue[0]
			nodeQueue = nodeQueue[1:]

			layoutList = append(layoutList, n.Val)
			if n.Left != nil {
				nodeQueue = append(nodeQueue, n.Left)
			}
			if n.Right != nil {
				nodeQueue = append(nodeQueue, n.Right)
			}
		}
		ret = append(ret, layoutList)
	}
	return ret
}

func main() {
	node := &treenode.TreeNode{
		Val:  3,
		Left: &treenode.TreeNode{Val: 9},
		Right: &treenode.TreeNode{
			Val:   20,
			Left:  &treenode.TreeNode{Val: 15},
			Right: &treenode.TreeNode{Val: 7},
		},
	}
	expectRet := [][]int{
		{3}, {9, 20}, {15, 7},
	}

	ret := LevelOrderByBFS(node)
	if !reflect.DeepEqual(ret, expectRet) {
		panic(fmt.Errorf("code102 error, LevelOrderByBFS() result is not expect result: \n%+v", ret))
	}

	ret = LevelOrderByDFS(node)
	if !reflect.DeepEqual(ret, expectRet) {
		panic(fmt.Errorf("code102 error, LevelOrderByDFS() result is not expect result: \n%+v", ret))
	}
}
