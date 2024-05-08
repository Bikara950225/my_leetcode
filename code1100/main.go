package main

func numKLenSubstrNoRepeats(s string, k int) (ret int) {
	m := map[byte]bool{}
	si, qi := 0, 0

	for qi < len(s) {
		item := s[qi]
		l := qi - si + 1

		if m[item] || l > k {
			delete(m, s[si])
			si++
		} else {
			m[item] = true
			if l == k {
				ret++
			}
			qi++
		}
	}
	return
}

func main() {

}
