package main

import (
	"fmt"
	"math/rand"
)

type entry struct {
	value int
	count int
}

func sort(eList []*entry) {
	var sortFunc func(eList []*entry, b, e int)
	sortFunc = func(eList []*entry, b, e int) {
		if b >= e {
			return
		}

		rIdx := b + (rand.Int() % (e - b + 1))
		mark := eList[rIdx]
		eList[e], eList[rIdx] = eList[rIdx], eList[e]

		i, j := b-1, b
		for ; j < e; j++ {
			if eList[j].count > mark.count {
				i++
				eList[i], eList[j] = eList[j], eList[i]
			}
		}

		i++
		eList[i], eList[e] = eList[e], eList[i]

		sortFunc(eList, b, i-1)
		sortFunc(eList, i+1, e)
	}
	sortFunc(eList, 0, len(eList)-1)
}

func topKFrequent(nums []int, k int) (ret []int) {
	eMap := map[int]*entry{}
	for _, num := range nums {
		if _, ok := eMap[num]; !ok {
			eMap[num] = &entry{value: num}
		}
		eMap[num].count++
	}

	eList := make([]*entry, 0, len(eMap))
	for _, v := range eMap {
		eList = append(eList, v)
	}

	sort(eList)

	for i := 0; i < k; i++ {
		ret = append(ret, eList[i].value)
	}
	return
}

func main() {
	ret := topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
	fmt.Println(ret)
}
