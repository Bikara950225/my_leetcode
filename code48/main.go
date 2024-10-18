package main

func rotate(matrix [][]int) {
	rotate1(matrix)
}

func rotate1(matrix [][]int) {
	l := len(matrix)
	cache := make([][]int, l)
	for i := range cache {
		cache[i] = make([]int, l)
	}

	for i := range matrix {
		cachej := l - i - 1
		cachei := 0

		for j := range matrix[i] {
			cache[cachei][cachej] = matrix[i][j]
			cachei++
		}
	}

	for i := range cache {
		for j := range cache[i] {
			matrix[i][j] = cache[i][j]
		}
	}
}

func rotate2(matrix [][]int) {
	// 先进行行置换
	i, j := 0, len(matrix)-1
	for i < j {
		for r := 0; r < len(matrix); r++ {
			matrix[i][r], matrix[j][r] = matrix[j][r], matrix[i][r]
		}
		i++
		j--
	}
	// 再进行对折
	for i := range matrix {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

func main() {

}
