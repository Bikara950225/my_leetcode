package main

import "fmt"

func scanGrid(grid [][]byte, gridMark [][]bool, y, x int) {
	yLen := len(grid)
	xLen := len(grid[0])
	if y < 0 || y >= yLen {
		return
	}
	if x < 0 || x >= xLen {
		return
	}
	if grid[y][x] == '0' || gridMark[y][x] {
		return
	}

	gridMark[y][x] = true

	scanGrid(grid, gridMark, y-1, x) // 上
	scanGrid(grid, gridMark, y+1, x) // 下
	scanGrid(grid, gridMark, y, x-1) // 左
	scanGrid(grid, gridMark, y, x+1) // 右
}

func numIslands(grid [][]byte) int {
	ret := 0

	yLen := len(grid)
	xLen := len(grid[0])

	gridMark := make([][]bool, yLen)
	for i := range gridMark {
		gridMark[i] = make([]bool, xLen)
	}

	for y := 0; y < yLen; y++ {
		for x := 0; x < xLen; x++ {
			if grid[y][x] == '1' && !gridMark[y][x] {
				ret += 1
				scanGrid(grid, gridMark, y, x)
			}
		}
	}
	return ret
}

func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}

	ret := numIslands(grid)
	expectRet := 1
	if ret != expectRet {
		panic(fmt.Errorf("code200 error, not expect result: %v", ret))
	}
}
