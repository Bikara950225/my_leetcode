package tools

func IntSwap(src1, src2 *int) {
	if src1 == src2 || src1 == nil || src2 == nil {
		return
	}
	*src1 ^= *src2
	*src2 ^= *src1
	*src1 ^= *src2
}
