package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) (result string) {
	for i := 0; ; i++ {
		var c byte
		for j, str := range strs {
			if i >= len(str) {
				goto ret
			}
			if j == 0 {
				c = str[i]
			} else {
				if str[i] != c {
					goto ret
				}
			}
		}
		result += string(c)
	}

ret:
	return
}

func longestCommonPrefix2(strs []string) string {
	sb := strings.Builder{}

	i := 0
	for {
		var b byte
		for j, s := range strs {
			if i >= len(s) {
				goto exit
			}
			if j == 0 {
				b = s[i]
				continue
			}
			if b != s[i] {
				goto exit
			}
		}
		sb.WriteByte(strs[0][i])
		i++
	}
exit:
	return sb.String()
}

func main() {
	ret1 := longestCommonPrefix2([]string{"flower", "flow", "flight"})
	if ret1 != "fl" {
		panic(fmt.Sprintf(`code14 error, ret1 != "fl": %s`, ret1))
	}
}
