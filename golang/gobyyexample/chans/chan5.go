package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(time.Second * 1)
	boom := time.After(time.Second * 3)
	for {

		select {
		case <-tick:
			fmt.Println("滴答", time.Now().String())
		case <-boom:
			fmt.Println("碰")
		default:
			fmt.Println("吱吱吱吱...")
			time.Sleep(time.Second * 1)
		}
	}
}
