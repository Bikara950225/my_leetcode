package sort

import "math/rand"

func QuickSort(src []int) {
	quickSort2(src, 0, len(src)-1)
}

func quickSort1(nums []int, b, e int) {
	if b >= e {
		return
	}

	//randI := (rand.Int() % (e - b + 1)) + b
	//nums[randI], nums[e] = nums[e], nums[randI]
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

	// 随机抽样
	rIdx := b + rand.Intn(e-b+1)
	nums[e], nums[rIdx] = nums[rIdx], nums[e]
	m := nums[e]

	i, j := b, b-1
	ee := e
	for ; i < ee; i++ {
		if nums[i] < m {
			j++
			nums[i], nums[j] = nums[j], nums[i]
		} else if nums[i] == m {
			// 将与尾端相同的项放置到后面
			ee--
			nums[i], nums[ee] = nums[ee], nums[i]
		}
	}

	jj := j
	for ; ee <= e; ee++ {
		jj++
		nums[jj], nums[ee] = nums[ee], nums[jj]
	}

	quickSort2(nums, b, j)
	quickSort2(nums, jj+1, e)
}
