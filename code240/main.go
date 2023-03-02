package main

// 时间复杂度: M * logN
func method1(matrix [][]int, target int) bool {
	for _, sub := range matrix {
		if binarySearch(sub, target) {
			return true
		}
	}
	return false
}

func method2(matrix [][]int, target int) bool {
	beginY, beginX := 0, len(matrix[0])-1
	for beginY < len(matrix) && beginX >= 0 {
		currVal := matrix[beginY][beginX]
		if target == currVal {
			return true
		}

		if target < currVal {
			beginX -= 1
			continue
		}
		if target > currVal {
			beginY += 1
			continue
		}
	}
	return false
}

func searchMatrix(matrix [][]int, target int) bool {
	return method1(matrix, target)
}

func binarySearch(src []int, target int) bool {
	i, j := 0, len(src)-1
	for i < j {
		m := (i + j) >> 1
		if src[m] >= target {
			j = m
		} else {
			i = m + 1
		}
	}
	return src[i] == target
}

func main() {

}
