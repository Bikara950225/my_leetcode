package sort

import (
	"math/rand"
)

func QuickSort(src []int) {
	quickSort(src, 0, len(src)-1)
}

func quickSort(nums []int, b, e int) {
	if b >= e {
		return
	}

	randI := (rand.Int() % (e - b + 1)) + b
	mark := nums[randI]
	nums[randI], nums[e] = nums[e], nums[randI]

	i, j := b-1, b
	for ; j < e; j++ {
		if nums[j] < mark {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	i++
	nums[e], nums[i] = nums[i], nums[e]

	quickSort(nums, b, i-1)
	quickSort(nums, i+1, e)
}

//func quickSort(src []int, begin, end int) {
//	if begin >= end {
//		return
//	}
//
//	randI := (rand.Int() % (end - begin + 1)) + begin
//	mark := src[randI]
//	tools.IntSwap(&src[randI], &src[end])
//
//	i := begin - 1
//	for j := begin; j < end; j++ {
//		if src[j] < mark {
//			i++
//			tools.IntSwap(&src[i], &src[j])
//		}
//	}
//	i++
//	tools.IntSwap(&src[i], &src[end])
//
//	quickSort(src, begin, i-1)
//	quickSort(src, i+1, end)
//}
