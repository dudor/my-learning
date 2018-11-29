package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(time.Second * 3)
	fmt.Println(time.Now())
	<-timer1.C
	fmt.Println(time.Now())
	<-timer1.C
	fmt.Println(time.Now())

	//	time2 := time.NewTimer(time.Second * 2)
	//	go func() {
	//		<-time2.C
	//		fmt.Println(time.Now())
	//	}()

	//	fmt.Println(time.Now())
	//	<-time2.C
	//	fmt.Println(time.Now())
	//	<-time2.C
}
