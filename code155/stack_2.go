package main

// 不需要辅助栈
type MinStack2 struct {
	stack []int
	min   int
}

var _ Stack = (*MinStack2)(nil)

func NewMinStack2() MinStack2 {
	return MinStack2{
		stack: []int{},
		min:   0,
	}
}

func (s *MinStack2) Push(val int) {
	if len(s.stack) <= 0 {
		s.stack = append(s.stack, val)
		s.min = val
		return
	}
	compare := val - s.min
	s.stack = append(s.stack, compare)
	if compare < 0 {
		// compare < 0; 说明当前插入的值小于min，更新min
		s.min = val
	}
}

func (s *MinStack2) Pop() {
	last := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	if last < 0 {
		s.min = s.min - last
	}
}

func (s *MinStack2) Top() int {
	last := s.stack[len(s.stack)-1]

	return last + s.min
}

func (s *MinStack2) GetMin() int {
	return s.min
}
