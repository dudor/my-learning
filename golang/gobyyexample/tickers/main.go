package main

import (
	"fmt"
	"time"
)

func main() {
	ticker1 := time.NewTicker(time.Second * 2)
	go func() {
		for i := range ticker1.C {
			fmt.Println(time.Now(), i)
		}
	}()

	timer1 := time.NewTimer(time.Second * 6)
	<-timer1.C
	ticker1.Stop()
	time.Sleep(100 * time.Second)
}
