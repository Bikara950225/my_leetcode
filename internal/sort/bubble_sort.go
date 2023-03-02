package sort

import "my_leetcode/internal/tools"

func BubbleSort(src []int) {
	for i := 0; i < len(src); i++ {
		for j := 0; j < len(src)-1; j++ {
			if src[j] > src[j+1] {
				tools.IntSwap(&src[j], &src[j+1])
			}
		}
	}
}
