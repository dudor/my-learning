package main

import (
	"fmt"
	"time"
)

func main() {
	yzj := make(chan int, 5)
	go fibo(cap(yzj), yzj)
	value, ok := <-yzj
	fmt.Println(value, ok)
	for i := range yzj {
		fmt.Println(i)
	}
	value, ok = <-yzj
	fmt.Println(value, ok)
	time.Sleep(time.Second)
}

func fibo(num int, producer chan int) {
	x, y := 1, 1
	for i := 0; i < num; i++ {
		producer <- x
		x, y = y, x+y
	}
	close(producer)
}
