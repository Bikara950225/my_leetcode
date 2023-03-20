package main

import (
	"fmt"
	treenode "my_leetcode/internal/tree_node"
	"reflect"
)

func deleteNode(node *treenode.TreeNode, key int) *treenode.TreeNode {
	if node == nil {
		return nil
	}

	if node.Val < key {
		node.Right = deleteNode(node.Right, key)
	} else if node.Val > key {
		node.Left = deleteNode(node.Left, key)
	} else {
		// node.Val == key
		// 找到左子树的最大节点（即左子树的最右节点）,将其替换到当前替换节点，指针可以不变，但是其Val和子指针需要变化
		if node.Left == nil {
			return node.Right
		}
		var maxInt int
		var deleteFunc func(*treenode.TreeNode) *treenode.TreeNode
		deleteFunc = func(n *treenode.TreeNode) *treenode.TreeNode {
			if n.Right == nil {
				maxInt = n.Val
				return n.Left
			}
			n.Right = deleteFunc(n.Right)
			return n
		}
		node.Left = deleteFunc(node.Left)
		node.Val = maxInt
	}
	return node
}

func main() {
	ret1 := deleteNode(&treenode.TreeNode{Val: 5}, 5)
	if ret1 != nil {
		panic(fmt.Errorf("code450 error, not expect nil: %v", ret1))
	}

	ret2 := deleteNode(&treenode.TreeNode{
		Val: 4,
		Right: &treenode.TreeNode{
			Val: 7,
			Left: &treenode.TreeNode{
				Val: 6,
				Left: &treenode.TreeNode{
					Val: 5,
				},
			},
			Right: &treenode.TreeNode{
				Val: 8,
				Right: &treenode.TreeNode{
					Val: 9,
				},
			},
		},
	}, 7)
	expectRet2 := &treenode.TreeNode{
		Val: 4,
		Right: &treenode.TreeNode{
			Val: 6,
			Left: &treenode.TreeNode{
				Val: 5,
			},
			Right: &treenode.TreeNode{
				Val: 8,
				Right: &treenode.TreeNode{
					Val: 9,
				},
			},
		},
	}
	if !reflect.DeepEqual(ret2, expectRet2) {
		panic(fmt.Errorf("code450 error, ret2: %v", ret2))
	}
}
