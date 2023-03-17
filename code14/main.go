package main

import "fmt"

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

func main() {
	ret1 := longestCommonPrefix([]string{"flower", "flow", "flight"})
	if ret1 != "fl" {
		panic(fmt.Sprintf(`code14 error, ret1 != "fl": %s`, ret1))
	}
}
