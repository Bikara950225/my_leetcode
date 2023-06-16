package main

import "fmt"

func toStr(src []int) (ret string) {
	for _, item := range src {
		ret += fmt.Sprint(item)
	}
	return
}

func permuteUnique(nums []int) (ret [][]int) {
	numSet := map[int]struct{}{}
	retSet := map[string]struct{}{}
	var subRet []int

	var parseFunc func()
	parseFunc = func() {
		if len(subRet) == len(nums) {
			subRetStr := toStr(subRet)
			if _, ok := retSet[subRetStr]; !ok {
				retSet[subRetStr] = struct{}{}

				cpRet := make([]int, len(subRet))
				copy(cpRet, subRet)
				ret = append(ret, cpRet)
				return
			}
		}

		for i, item := range nums {
			if _, ok := numSet[i]; ok {
				continue
			}

			numSet[i] = struct{}{}
			subRet = append(subRet, item)

			parseFunc()

			delete(numSet, i)
			subRet = subRet[:len(subRet)-1]
		}
	}

	parseFunc()
	return
}

func main() {

}
