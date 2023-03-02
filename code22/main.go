package main

import (
	"fmt"
	"reflect"
)

type solution struct {
	ret      []string
	strCache []byte
}

func (s *solution) dfs(n int, leftN, rightN int) {
	if len(s.strCache) == n*2 {
		s.ret = append(s.ret, string(s.strCache))
		return
	}

	if leftN < n {
		s.strCache = append(s.strCache, '(')
		s.dfs(n, leftN+1, rightN)
		s.strCache = s.strCache[:len(s.strCache)-1]
	}
	if rightN < leftN {
		s.strCache = append(s.strCache, ')')
		s.dfs(n, leftN, rightN+1)
		s.strCache = s.strCache[:len(s.strCache)-1]
	}
}

func (s *solution) GenerateParenthesis(n int) []string {
	s.dfs(n, 0, 0)
	return s.ret
}

func main() {
	ss := &solution{}
	ret := ss.GenerateParenthesis(3)
	expectRet := []string{
		"((()))", "(()())", "(())()", "()(())", "()()()",
	}
	if !reflect.DeepEqual(expectRet, ret) {
		panic(fmt.Errorf("GenerateParenthesis() error, isn't expectRet: \n%+v", ret))
	}
}
