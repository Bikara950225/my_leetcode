package sort

import "math/rand"

func QuickSort(src []int) {
	quickSort2(src, 0, len(src)-1)
}

func quickSort1(nums []int, b, e int) {
	if b >= e {
		return
	}

	randI := (rand.Int() % (e - b + 1)) + b
	nums[randI], nums[e] = nums[e], nums[randI]
	// 这里不随机抽样
	mark := nums[e]

	i, j := b-1, b
	for ; j < e; j++ {
		if nums[j] < mark {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	i++
	nums[e], nums[i] = nums[i], nums[e]

	quickSort1(nums, b, i-1)
	quickSort1(nums, i+1, e)
}

func quickSort2(nums []int, b, e int) {
	if b >= e {
		return
	}

	idx := (rand.Int() % (e - b + 1)) + b
	mark := nums[idx]
	nums[idx], nums[e] = nums[e], nums[idx]

	i, j := b, b-1
	for ; i < e; i++ {
		if nums[i] <= mark {
			j++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	j++
	nums[j], nums[e] = nums[e], nums[j]

	quickSort2(nums, b, j-1)
	quickSort2(nums, j+1, e)
}
