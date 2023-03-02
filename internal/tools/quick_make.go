package tools

// 快速生成 leetcode 中使用的二维图、DP等数据结构

// 快速创建二维的DP数组
// 创建后二维数组的每一个元素都为0值
func NewInt2DP(x, y int) [][]int {
	ret := make([][]int, 0, y)
	for i := 0; i < y; i++ {
		ret = append(ret, make([]int, x))
	}
	return ret
}

// 生成元素为byte的二维图
func NewByteMatrix(lines ...[]byte) [][]byte {
	matrix := make([][]byte, 0, len(lines))
	for _, line := range lines {
		matrix = append(matrix, line)
	}
	return matrix
}
