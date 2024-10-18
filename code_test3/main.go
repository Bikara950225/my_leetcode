package main

import (
	"fmt"
	"html/template"
	"os"
)

//
//func method(src []int) (ret int) {
//	for i, num := range src {
//		subRet := 0
//		// <-
//		for j := i - 1; j >= 0; j-- {
//
//		}
//		// ->
//		for j := i + 1; j < len(src); j++ {
//
//		}
//
//		ret = max(ret, subRet)
//	}
//	return
//}

func doubleG() {
	ch := make(chan struct{})

	go func() {
		for range 50 {
			<-ch
			fmt.Println("A")
			ch <- struct{}{}
		}
		if _, ok := <-ch; ok {
			close(ch)
		}
	}()

	ch <- struct{}{}
	for range 50 {
		<-ch
		fmt.Println("B")
		ch <- struct{}{}
	}
	if _, ok := <-ch; ok {
		close(ch)
	}
}

func main() {
	l := []string{"1", "2"}
	src := `key = [{{range $idx, $item := .}}{{if $idx}},{{end}}"{{$item}}"{{end}}]`

	tt, err := template.New("123").Parse(src)
	if err != nil {
		panic(err)
	}
	tt.Execute(os.Stdout, l)
}
