package main

import (
	"fmt"
)

func main() {
	chan_name := make(chan int)
	chan_quit := make(chan int)

	go func() {
		for i := 0; i < 20; i++ {
			fmt.Println(<-chan_name)
		}
		chan_quit <- 100
	}()
	fibo(chan_name, chan_quit)
}

func fibo(chan_name, chan_quit chan int) {
	x, y := 0, 1
	for {
		select {
		case chan_name <- x:
			x, y = y, x+y
		case <-chan_quit:
			fmt.Println("EXIT")
			return
		}
	}
}
