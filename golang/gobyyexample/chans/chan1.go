package main

import "fmt"

func main() {
	chan1 := make(chan string)
	go func() {
		data := <-chan1
		fmt.Println(data)

	}()
	chan1 <- "job1"

	fmt.Println("done")
}
