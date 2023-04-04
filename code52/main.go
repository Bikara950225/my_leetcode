package main

import "fmt"

func totalNQueens(n int) (ret int) {
	var xLine int
	yMap := map[int]struct{}{}
	mat := make([][]int, n)
	for i := range mat {
		mat[i] = make([]int, n)
	}

	checkFunc := func(y int) bool {
		if _, ok := yMap[y]; ok {
			return false
		}
		by := y

		x := xLine
		for x >= 0 && y >= 0 {
			if mat[x][y] == 1 {
				return false
			}
			x--
			y--
		}

		x, y = xLine, by
		for x >= 0 && y < n {
			if mat[x][y] == 1 {
				return false
			}
			x--
			y++
		}
		return true
	}

	var dfsFunc func()
	dfsFunc = func() {
		if xLine == n {
			ret += 1
			return
		}

		for i := 0; i < n; i++ {
			if !checkFunc(i) {
				continue
			}

			yMap[i] = struct{}{}
			mat[xLine][i] = 1
			xLine += 1

			dfsFunc()

			xLine -= 1
			mat[xLine][i] = 0
			delete(yMap, i)
		}
	}
	dfsFunc()

	return
}

func main() {
	ret1 := totalNQueens(4)
	if ret1 != 2 {
		panic(fmt.Errorf("code52 error, ret1 != 2: %d", ret1))
	}
}
