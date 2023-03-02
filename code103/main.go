package main

import (
	"fmt"
	internaltree "my_leetcode/internal/tree_node"
	"reflect"
)

func zigzagLevelOrder(root *internaltree.TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var ret [][]int
	nodeQueue := []*internaltree.TreeNode{root}
	level := 0
	for len(nodeQueue) > 0 {
		layoutSize := len(nodeQueue)
		layout := make([]int, layoutSize)
		ret = append(ret, layout)
		for i := 0; i < layoutSize; i++ {
			curr := nodeQueue[0]
			nodeQueue = nodeQueue[1:]
			if (level & 0x1) == 0x1 {
				layout[layoutSize-i-1] = curr.Val
			} else {
				layout[i] = curr.Val
			}
			if curr.Left != nil {
				nodeQueue = append(nodeQueue, curr.Left)
			}
			if curr.Right != nil {
				nodeQueue = append(nodeQueue, curr.Right)
			}
		}
		level++
	}
	return ret
}

func main() {
	// test tree
	tree := &internaltree.TreeNode{
		Val:  3,
		Left: &internaltree.TreeNode{Val: 9},
		Right: &internaltree.TreeNode{
			Val:   20,
			Left:  &internaltree.TreeNode{Val: 15},
			Right: &internaltree.TreeNode{Val: 7},
		},
	}

	ret := zigzagLevelOrder(tree)
	expectRet := [][]int{
		{3}, {20, 9}, {15, 7},
	}

	if !reflect.DeepEqual(ret, expectRet) {
		panic(fmt.Errorf("not expect result: \n%+v", ret))
	}
}
