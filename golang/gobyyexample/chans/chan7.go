package main

import (
	"fmt"
	"time"
)

func main() {
	data := make(chan int)
	go Parse(data)
	data <- 1000
	time.Sleep(time.Second)
}

func Parse(ch <-chan int) {
	for k := range ch {
		fmt.Println(k)
	}
}
