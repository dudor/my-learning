package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 34, 1, 9, 32, 4, 33, 22, 11}
	sum := make(chan int)
	go Sum(arr[:len(arr)/2], sum)
	go Sum(arr[len(arr)/2:], sum)

	r1, r2 := <-sum, <-sum
	fmt.Println("sum = ", r1+r2)
}

func Sum(a []int, sum chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	sum <- total
}
