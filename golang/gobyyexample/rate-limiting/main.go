package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		requests <- i
	}
	close(requests)

	ticker1 := time.Tick(time.Second * 1)

	for req := range requests {
		<-ticker1
		fmt.Println(req)

	}

}
