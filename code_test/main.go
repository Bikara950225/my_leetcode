package main

import (
	"encoding/json"
	"math"
	"math/rand"
	treenode "my_leetcode/internal/tree_node"
	"reflect"
	"sync"
	"unsafe"
)

const (
	e1 = iota
	e2
	e3
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

func unsafeStringToBytes(src string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&src))
	println(sh.Data)

	return (*[math.MaxInt32]byte)(unsafe.Pointer(sh.Data))[:sh.Len:sh.Len]
}

func unsafeBytesToString(src []byte) string {
	//return *(*string)(unsafe.Pointer(&src))
	sliH := (*reflect.SliceHeader)(unsafe.Pointer(&src))
	println(sliH.Data)
	return *(*string)(unsafe.Pointer(&src))
}

func quickSort(src []int, b, e int) {
	if b >= e {
		return
	}

	randIdx := b + (rand.Int() % (e - b + 1))
	mark := src[randIdx]
	src[e], src[randIdx] = src[randIdx], src[e]

	i, j := b-1, b
	for ; j < e; j++ {
		if src[j] < mark {
			i++
			src[i], src[j] = src[j], src[i]
		}
	}
	i++
	src[i], src[e] = src[e], src[i]

	quickSort(src, b, i-1)
	quickSort(src, i+1, e)
}

type ss string

type ssi interface {
	string | ss | int
	comparable
}

func method[T ~string](src1, src2 T) bool {
	return src1 == src2
}

type SyncMap[K comparable, V any] struct {
	mutex sync.Mutex
	cache map[K]V
}

func NewSyncMap[K1 comparable, V1 any]() *SyncMap[K1, V1] {
	return &SyncMap[K1, V1]{
		cache: make(map[K1]V1),
	}
}

func (s *SyncMap[K, V]) Get(key K) (V, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	v, ok := s.cache[key]
	return v, ok
}

func main() {
	sm := NewSyncMap[string, int]()
	sm.Get("123")
}
