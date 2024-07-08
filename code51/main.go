package main

import (
	"fmt"
	"strings"
)

func solveNQueens(n int) (ret [][]string) {
	cache := make([][]bool, n)
	for i := range cache {
		cache[i] = make([]bool, n)
	}

	ymap := make([]bool, n)
	var handleFunc func(x int)
	handleFunc = func(x int) {
		if x >= n {
			subRet := make([]string, 0, n)
			for _, subl := range cache {
				subStr := strings.Builder{}
				for _, item := range subl {
					if item {
						subStr.WriteString("Q")
					} else {
						subStr.WriteString(".")
					}
				}
				subRet = append(subRet, subStr.String())
			}
			ret = append(ret, subRet)
			return
		}

		for y := 0; y < n; y++ {
			if ymap[y] {
				continue
			}

			if !checkCache(cache, x, y) {
				continue
			}

			cache[x][y] = true
			ymap[y] = true
			handleFunc(x + 1)
			cache[x][y] = false
			ymap[y] = false
		}
	}

	handleFunc(0)
	return
}

func checkCache(cache [][]bool, x, y int) bool {
	xx, yy := x-1, y-1
	for xx >= 0 && yy >= 0 {
		if cache[xx][yy] {
			return false
		}
		xx--
		yy--
	}

	xx, yy = x-1, y+1
	for xx >= 0 && yy < len(cache) {
		if cache[xx][yy] {
			return false
		}
		xx--
		yy++
	}
	return true
}

func main() {
	ret1 := solveNQueens(5)
	fmt.Println(ret1)
}
