package main

import (
	"fmt"
)

func main() {
	go func() {
		fmt.Println("sdfsdf")
	}()
	q1 := make(chan string, 2)
	q1 <- "n1"
	q1 <- "n2"
	close(q1)

	for a := range q1 {
		fmt.Println(a)
	}

}
