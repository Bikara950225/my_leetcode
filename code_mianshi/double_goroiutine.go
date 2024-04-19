package code_mianshi

import "fmt"

func method() {
	ch := make(chan struct{})

	fn := func(r string, n int) {
		for i := 0; i < n; i++ {
			<-ch
			fmt.Println(r)
			ch <- struct{}{}
		}
		if _, ok := <-ch; ok {
			close(ch)
		}
	}
	go fn("A", 50)
	ch <- struct{}{}
	fn("B", 50)
}
