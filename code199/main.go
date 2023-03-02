package main

import (
	"fmt"
	treenode "my_leetcode/internal/tree_node"
	"reflect"
)

func rightSideView(root *treenode.TreeNode) []int {
	if root == nil {
		return nil
	}

	var ret []int
	nodeQueue := []*treenode.TreeNode{
		root,
	}
	for len(nodeQueue) > 0 {
		layoutSize := len(nodeQueue)
		var targetNode *treenode.TreeNode
		for i := 0; i < layoutSize; i++ {
			targetNode = nodeQueue[0]
			nodeQueue = nodeQueue[1:]

			if targetNode.Left != nil {
				nodeQueue = append(nodeQueue, targetNode.Left)
			}
			if targetNode.Right != nil {
				nodeQueue = append(nodeQueue, targetNode.Right)
			}
		}
		if targetNode != nil {
			ret = append(ret, targetNode.Val)
		}
	}
	return ret
}

func main() {
	srcTree := &treenode.TreeNode{
		Val: 1,
		Left: &treenode.TreeNode{
			Val: 2,
			Right: &treenode.TreeNode{
				Val: 5,
			},
		},
		Right: &treenode.TreeNode{
			Val: 3,
			Right: &treenode.TreeNode{
				Val: 4,
			},
		},
	}
	ret := rightSideView(srcTree)
	expectRet := []int{1, 3, 4}
	if !reflect.DeepEqual(ret, expectRet) {
		panic(fmt.Errorf("code199 error, not expect result: %+v", ret))
	}
}
