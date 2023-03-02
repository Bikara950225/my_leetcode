package main

import (
	"fmt"
	"reflect"
)

var (
	ret     []string
	ipCache []int
)

func restoreIpAddresses(s string) []string {
	defer func() {
		ret = nil
		ipCache = nil
	}()
	dfs(s, 0)
	return ret
}

func dfs(s string, sIndex int) {
	if sIndex == len(s) {
		if len(ipCache) == 4 {
			ipStr := fmt.Sprintf("%d.%d.%d.%d", ipCache[0], ipCache[1], ipCache[2], ipCache[3])
			ret = append(ret, ipStr)
		}
		return
	}
	if len(ipCache) >= 4 {
		// ipCache大于等于4段时，sIndex还没到末尾，直接退出
		return
	}

	if s[sIndex] == '0' {
		ipCache = append(ipCache, 0)
		dfs(s, sIndex+1)
		ipCache = ipCache[:len(ipCache)-1]
		return
	}

	var subIp int
	for si := sIndex; si < len(s); si++ {
		// s[si] - '0' 是 字符数字 转换成 整形数字 的小方法
		subIp = (10 * subIp) + int(s[si]-'0')
		// ip子段不能超出 255（0xFF）
		if subIp > 0xFF {
			break
		}

		ipCache = append(ipCache, subIp)
		dfs(s, si+1)
		ipCache = ipCache[:len(ipCache)-1]
	}
}

func main() {
	s1 := "25525511135"
	ret1 := restoreIpAddresses(s1)
	expectRet1 := []string{"255.255.11.135", "255.255.111.35"}
	if !reflect.DeepEqual(ret1, expectRet1) {
		panic(fmt.Errorf("code93 error, not expect result: %+v", ret1))
	}
}
