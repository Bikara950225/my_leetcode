package main

import (
	"fmt"
	"my_leetcode/internal/sort"
)

func sortArray(nums []int) []int {
	sort.QuickSort(nums)
	//sort.MergeSort(nums)
	//sort.HeapSort(nums)

	return nums
}

func main() {
	ret := sortArray([]int{3, 2, 1})
	fmt.Println(ret)
}
