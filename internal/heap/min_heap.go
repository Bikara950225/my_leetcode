package heap

type MinHeap []int

func (s *MinHeap) Add(item int) {
	*s = append(*s, item)
	s.heapifyUp()
}

func (s *MinHeap) Top() (int, bool) {
	if len(*s) <= 0 {
		return -1, false
	}
	return (*s)[0], true
}

func (s *MinHeap) Pop() (int, bool) {
	top, ok := s.Top()
	if !ok {
		return -1, false
	}

	end := len(*s) - 1
	(*s)[0], (*s)[end] = (*s)[end], (*s)[0]
	*s = (*s)[:len(*s)-1]
	s.heapifyUp()

	return top, true
}

func (s *MinHeap) heapifyUp() {
	l := len(*s)
	for i := l >> 1; i >= 0; i-- {
		bg := i
		for bg<<1+1 < l {
			litter := bg
			// TODO 后续试试实现为四叉堆
			for _, sub := range [2]int{bg<<1 + 1, bg<<1 + 2} {
				if sub < l {
					if (*s)[litter] > (*s)[sub] {
						litter = sub
					}
				} else {
					break
				}
			}
			if bg != litter {
				(*s)[bg], (*s)[litter] = (*s)[litter], (*s)[bg]
				bg = litter
			} else {
				break
			}
		}
	}
}
