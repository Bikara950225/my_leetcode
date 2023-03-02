package main

type Stack interface {
	Push(val int)
	Pop()
	Top() int
	GetMin() int
}
