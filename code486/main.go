package main

import (
	"fmt"
	"unsafe"
)

type TestStruct1 struct{}
type TestStruct2 struct{}

func (s *TestStruct1) Method() {
	fmt.Println("TestStruct1")
}

func (s *TestStruct2) Method() {
	fmt.Println("TestStruct2")
}

func main() {
	var ts1 TestStruct1
	var ts2 TestStruct2

	ts1 = *(*TestStruct1)(unsafe.Pointer(&ts2))
	ts1.Method()
}
