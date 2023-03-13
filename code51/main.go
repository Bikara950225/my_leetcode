package main

import "fmt"

var (
	ret [][]string
	m   = map[int]struct{}{}
)

func solveNQueens(n int) [][]string {
	defer func() {
		ret = nil
		m = map[int]struct{}{}
	}()

	ml := make([][]int, n)
	for i := range ml {
		ml[i] = make([]int, n)
	}

	dfs(n, ml, 0)

	return ret
}

func dfs(n int, ml [][]int, x int) {
	if x == n {
		subRet := make([]string, 0, n)
		for _, xItem := range ml {
			var subStr string
			for _, yItem := range xItem {
				if yItem == 0 {
					subStr += "."
				} else {
					subStr += "Q"
				}
			}
			subRet = append(subRet, subStr)
		}
		ret = append(ret, subRet)
		return
	}

	xLine := ml[x]
	for i := 0; i < n; i++ {
		if _, ok := m[i]; ok {
			continue
		}
		if !check(ml, n, x, i) {
			continue
		}

		xLine[i] = 1
		m[i] = struct{}{}

		dfs(n, ml, x+1)

		xLine[i] = 0
		delete(m, i)
	}
	return
}

func check(ml [][]int, n, x, y int) bool {
	x1, y1 := x-1, y-1
	for x1 >= 0 && y1 >= 0 {
		if ml[x1][y1] == 1 {
			return false
		}
		x1--
		y1--
	}

	x2, y2 := x-1, y+1
	for x2 >= 0 && y2 < n {
		if ml[x2][y2] == 1 {
			return false
		}
		x2--
		y2++
	}

	return true
}

func main() {
	ret1 := solveNQueens(4)
	fmt.Println(ret1)
}
