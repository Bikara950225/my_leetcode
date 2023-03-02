package main

// 使用辅助栈
type MinStack1 struct {
	stack []int
	// 辅助栈
	// 这里的栈因为会有重复元素的情况，所以可以考虑数据压缩
	// 比如 -2, -2, -3, -1, -1, 压缩为(-2, 2), (-3, 1), (-1, 2)
	minStack []int
}

func ConstructorMinStack1() MinStack1 {
	return MinStack1{
		stack:    []int{},
		minStack: []int{},
	}
}

func (s *MinStack1) Push(val int) {
	s.stack = append(s.stack, val)

	// 辅助栈
	if len(s.minStack) <= 0 {
		s.minStack = append(s.minStack, val)
	} else {
		minStackLast := s.minStack[len(s.minStack)-1]
		if val < minStackLast {
			s.minStack = append(s.minStack, val)
		} else {
			s.minStack = append(s.minStack, minStackLast)
		}
	}
}

func (s *MinStack1) Pop() {
	s.stack = s.stack[:len(s.stack)-1]
	s.minStack = s.minStack[:len(s.minStack)-1]
}

func (s *MinStack1) Top() int {
	return s.stack[len(s.stack)-1]
}

func (s *MinStack1) GetMin() int {
	return s.minStack[len(s.minStack)-1]
}
