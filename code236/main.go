package main

import treenode "my_leetcode/internal/tree_node"

var ret *treenode.TreeNode

func find(node, p, q *treenode.TreeNode) bool {
	if node == nil {
		return false
	}

	// 递归，左子树是否有目标节点
	leftCheck := find(node.Left, p, q)
	// 递归，右子树是否有目标节点
	rightCheck := find(node.Right, p, q)

	// 1. node的左、右节点都发现了p、q时
	// 2. node自身等于p或者q时，左节点发现了另一个节点时
	// 3. node自身等于p或者q时，右节点发现了另一个节点时
	pEqual := node.Val == p.Val
	qEqual := node.Val == q.Val
	currEqual := pEqual || qEqual

	if (leftCheck && rightCheck) || (currEqual && leftCheck) || (currEqual && rightCheck) {
		ret = node
	}
	return currEqual || leftCheck || rightCheck
}

func lowestCommonAncestor(root, p, q *treenode.TreeNode) *treenode.TreeNode {
	defer func(){
		ret = nil
	}()
	find(root, p, q)
	return ret
}

func main() {

}
