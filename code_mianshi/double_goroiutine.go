package code_mianshi

import (
	"fmt"
)

func doubleG() {
	ch := make(chan struct{})

	var ff = func(n int, content string) {
		for range n {
			<-ch
			fmt.Println(content)
			ch <- struct{}{}
		}
		if _, ok := <-ch; ok {
			close(ch)
		}
	}

	go ff(50, "A")

	ch <- struct{}{}
	ff(50, "B")
}

func doubleG2() {
	ch1 := make(chan struct{}, 1)
	ch2 := make(chan struct{}, 1)

	go func() {
		for i := range 50 {
			<-ch1
			fmt.Printf("[%d]A\n", i)
			ch2 <- struct{}{}
		}
	}()

	ch1 <- struct{}{}
	for i := range 50 {
		<-ch2
		fmt.Printf("[%d]B\n", i)
		ch1 <- struct{}{}
	}
}

func doubleGT() {
	n := 50
	ch := make(chan string)

	ff := func(s string) {
		for range n {
			fmt.Println(<-ch)
			ch <- s
		}
		if _, ok := <-ch; ok {
			close(ch)
		}
	}

	go ff("B")
	ch <- "A"
	ff("A")
}
