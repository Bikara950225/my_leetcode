package code_mianshi

import (
	"fmt"
	"github.com/emirpasic/gods/v2/maps/treemap"
	"github.com/google/btree"
	"sort"
)

// TreeMap通常使用平衡二叉树（多数为RB-Tree）实现，可以做到以O(logN)的时间复杂度进行查询，可以按照key的顺序遍历的效果；
// BTreeMap通常使用B树或者B+树构成，可以保证O(logN)的时间复杂度进行读取，且维护key的顺序遍历；相较于TreeMap，BtreeMap可以找到数据节点后，通过Iter迭代器进行遍历，因为所用数据结构会维护一个底层的双端链表；

type entry struct {
	id int
	v  string
}

func treeMapDemo() {
	tmap := treemap.New[int, *entry]()
	for i := range 1000 {
		tmap.Put(i, &entry{
			id: i, v: fmt.Sprintf("val%d", i),
		})
	}
	vl := tmap.Values()
	isSorted := sort.SliceIsSorted(vl, func(i, j int) bool {
		return vl[i].id < vl[j].id
	})
	fmt.Println(isSorted)
}

func bTreeMapDemo() {
	btmap := btree.NewG[*entry](10, func(a, b *entry) bool {
		return a.id < b.id
	})
	for i := range 1000 {
		btmap.ReplaceOrInsert(&entry{
			id: i, v: fmt.Sprintf("value%d", i),
		})
	}
	// 范围查询, [49, 101)
	btmap.AscendRange(&entry{id: 49}, &entry{id: 101}, func(item *entry) bool {
		fmt.Println(item)
		return true
	})
}

// 暂时没法find临近的左值后，通过iter遍历找到查询目标，达到稀疏索引的效果
func bTreeMapSparseIndexDemo() {
	n := 100000
	mainData := make([]*entry, n)
	for i := range n {
		mainData[i] = &entry{
			id: i, v: fmt.Sprintf("val%d", i),
		}
	}
	// 每100条建立一个索引
	sparseIdx := btree.NewG(n/100, func(a, b *entry) bool {
		return a.id < b.id
	})
	for i, e := range mainData {
		if i%100 == 0 {
			sparseIdx.ReplaceOrInsert(e)
		}
	}

	find := 549
	//sparseIdx.AscendGreaterOrEqual(findEntry, func(item *entry) bool {
	//	fmt.Println(item)
	//	return true
	//})
	var targetStart *entry
	sparseIdx.AscendLessThan(&entry{id: find + 1}, func(item *entry) bool {
		targetStart = item
		return true
	})

	idx := targetStart.id
	for ; idx <= find; idx++ {
		if mainData[idx].id == find {
			fmt.Println("find")
		}
	}
	fmt.Println("not found")
}
