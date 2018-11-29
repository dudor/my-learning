package main

import (
	"fmt"
)

func main() {
	data := make(chan int)

	go func() {
		write := chan<- int(data)
		write <- 100
		fmt.Println("write", write)
		fmt.Printf("write %T", write)
	}()

	read := <-chan int(data)
	fmt.Println("read", read)
	fmt.Printf("read %T", read)
}
