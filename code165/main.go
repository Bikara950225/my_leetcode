package main

import (
	"fmt"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	ver1List := strings.Split(version1, ".")
	ver2List := strings.Split(version2, ".")

	return dfs(ver1List, ver2List, 0, 0)
}

// 递归的话，当其中一个semversion有子版本号缺失，当0处理即可
// 不递归的话，可以比较2个semversion的版本长度，少的那方补0，然后再处理也可以~
func dfs(ver1List, ver2List []string, ver1I, ver2I int) int {
	if ver1I >= len(ver1List) && ver2I >= len(ver2List) {
		return 0
	}

	var subVer1 int
	if ver1I < len(ver1List) {
		subVer1 = parseSemVersion(ver1List[ver1I])
	}

	var subVer2 int
	if ver2I < len(ver2List) {
		subVer2 = parseSemVersion(ver2List[ver2I])
	}

	if subVer1 > subVer2 {
		return 1
	} else if subVer1 < subVer2 {
		return -1
	}
	return dfs(ver1List, ver2List, ver1I+1, ver2I+1)
}

func parseSemVersion(subVer string) int {
	verInt := 0
	for _, b := range subVer {
		verInt = (verInt * 10) + int(b-'0')
	}
	return verInt
}

func main() {
	ret1 := compareVersion("1.0.0", "1")
	if ret1 != 0 {
		panic(fmt.Errorf("code165 error, 1.0.0 compare 1, not expect: %d", ret1))
	}

	ret2 := compareVersion("1", "1.0.1")
	if ret2 != -1 {
		panic(fmt.Errorf("code165 error, 1 compare 1.0.1, not expect: %d", ret2))
	}

	ret3 := compareVersion("1.0.2", "1.0.1")
	if ret3 != 1 {
		panic(fmt.Errorf("code165 error, 1.0.2 compare 1.0.1, not expect: %d", ret3))
	}

	ret4 := compareVersion("2", "1.0.1")
	if ret4 != 1 {
		panic(fmt.Errorf("code165 error, 2 compare 1.0.1, not expect: %d", ret4))
	}
}
