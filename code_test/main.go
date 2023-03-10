package main

import "time"

func main() {
	ch := make(chan int, 1)
	for {
		select {
		case ch <- 1:
			println("ping")
			time.Sleep(time.Second)
		case <-ch:
			println("pong")
			time.Sleep(time.Second)
		}
	}

}
