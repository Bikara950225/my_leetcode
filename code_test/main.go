package main

import (
	"encoding/json"
	"fmt"
	"math"
	treenode "my_leetcode/internal/tree_node"
	"reflect"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
// 前缀遍历
func (s *Codec) serialize(root *treenode.TreeNode) string {
	if root == nil {
		return ""
	}
	var vList []int
	var q []*treenode.TreeNode
	for len(q) > 0 || root != nil {
		for root != nil {
			vList = append(vList, root.Val)
			q = append(q, root)
			root = root.Left
		}
		root = q[len(q)-1]
		q = q[:len(q)-1]
		root = root.Right
	}
	ret, _ := json.Marshal(vList)
	return string(ret)
}

// Deserializes your encoded data to tree.
func (s *Codec) deserialize(data string) *treenode.TreeNode {
	var dataList []int
	_ = json.Unmarshal([]byte(data), &dataList)

	var dfsFunc func(min, max int) *treenode.TreeNode
	dfsFunc = func(min, max int) *treenode.TreeNode {
		if len(dataList) <= 0 {
			return nil
		}
		currV := dataList[0]
		if min >= currV || max <= currV {
			return nil
		}
		dataList = dataList[1:]

		return &treenode.TreeNode{
			Val:   currV,
			Left:  dfsFunc(min, currV),
			Right: dfsFunc(currV, max),
		}
	}
	return dfsFunc(math.MinInt, math.MaxInt)
}

func main() {
	c := Codec{}

	ret1Tree := &treenode.TreeNode{
		Val: 3,
		Left: &treenode.TreeNode{
			Val: 1,
			Right: &treenode.TreeNode{
				Val: 2,
			},
		},
	}
	ret1_1 := c.serialize(ret1Tree)
	if ret1_1 != `[3,1,2]` {
		panic(fmt.Errorf("ret1_1 not match expect: %s", ret1_1))
	}
	ret1_2 := c.deserialize(ret1_1)
	if !reflect.DeepEqual(ret1Tree, ret1_2) {
		panic(fmt.Errorf("ret1_2 not match expect: %v", ret1_2))
	}
}
