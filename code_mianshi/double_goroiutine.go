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

func method2() {
	ch1 := make(chan struct{}, 1)
	ch2 := make(chan struct{}, 1)

	go func(n int) {
		for i := 0; i < n; i++ {
			<-ch1
			fmt.Println("A")
			ch2 <- struct{}{}
		}
	}(50)

	ch1 <- struct{}{}
	for i := 0; i < 50; i++ {
		<-ch2
		fmt.Println("B")
		ch1 <- struct{}{}
	}
}
