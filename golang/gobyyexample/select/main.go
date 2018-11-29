package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 3)
		c1 <- "test1"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "test2"
	}()
	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		case <-time.After(time.Second * 2):
			fmt.Println("timeout")
		}
	}
}
