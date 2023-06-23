package tools

func StringReverse(src string) string {
	srcR := []rune(src)
	i, j := 0, len(srcR)-1
	for i < j {
		srcR[i], srcR[j] = srcR[j], srcR[i]
		i++
		j--
	}
	return string(srcR)
}
