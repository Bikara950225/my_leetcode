package tree_node

import internalconst "my_leetcode/internal/consts"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (s *TreeNode) GetPreorder() []int {
	var ret []int
	var nodeStack []*TreeNode
	curr := s
	for curr != nil || len(nodeStack) > 0 {
		for curr != nil {
			ret = append(ret, curr.Val)
			nodeStack = append(nodeStack, curr)
			curr = curr.Left
		}
		// 出栈
		curr = nodeStack[len(nodeStack)-1]
		nodeStack = nodeStack[:len(nodeStack)-1]
		curr = curr.Right
	}
	return ret
}

func (s *TreeNode) GetInorder() []int {
	var ret []int
	var nodeStack []*TreeNode
	curr := s
	for curr != nil || len(nodeStack) > 0 {
		for curr != nil {
			nodeStack = append(nodeStack, curr)
			curr = curr.Left
		}
		// 出栈
		curr = nodeStack[len(nodeStack)-1]
		nodeStack = nodeStack[:len(nodeStack)-1]
		ret = append(ret, curr.Val)
		curr = curr.Right
	}
	return ret
}

func (s *TreeNode) GetPostorder() []int {
	var ret []int
	var nodeStack []*TreeNode
	curr := s
	var repeatNode *TreeNode

	for len(nodeStack) > 0 || curr != nil {
		for curr != nil {
			nodeStack = append(nodeStack, curr)
			curr = curr.Left
		}

		curr = nodeStack[len(nodeStack)-1]
		if curr.Right != nil && curr.Right != repeatNode {
			curr = curr.Right
		} else {
			nodeStack = nodeStack[:len(nodeStack)-1]
			ret = append(ret, curr.Val)
			repeatNode = curr
			curr = nil
		}
	}
	return ret
}

// 通过前序遍历构建二叉搜索树
// 要求输入的数组是一个严格的二叉搜索树的前缀遍历
func BuildSearchTreeWithPreorder(src []int) *TreeNode {
	var dfs func(lower, upper int) *TreeNode
	dfs = func(lower, upper int) *TreeNode {
		if len(src) <= 0 {
			return nil
		}

		currVal := src[0]
		if currVal < lower || currVal > upper {
			return nil
		}
		src = src[1:]

		return &TreeNode{
			Val:   currVal,
			Left:  dfs(lower, currVal),
			Right: dfs(currVal, upper),
		}
	}
	return dfs(int(internalconst.MinInt64), int(internalconst.MaxInt64))
}
