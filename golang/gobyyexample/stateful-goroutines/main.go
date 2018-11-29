package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key   int
	value int
	resp  chan bool
}

func main() {

	var readOps, writeOps int32

	readChan := make(chan *readOp)
	writeChan := make(chan *writeOp)

	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-readChan:
				read.resp <- state[read.key]

			case write := <-writeChan:
				state[write.key] = write.value
				write.resp <- true
			}

		}
	}()

	for i := 0; i < 100; i++ {
		go func() {
			for {
				r := &readOp{}
				r.key = rand.Intn(5)
				r.resp = make(chan int)

				readChan <- r
				<-r.resp
				atomic.AddInt32(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			for {
				w := &writeOp{
					key:   rand.Intn(5),
					value: rand.Intn(100),
					resp:  make(chan bool),
				}

				writeChan <- w
				<-w.resp
				atomic.AddInt32(&writeOps, 1)
				time.Sleep(time.Millisecond)

			}
		}()
	}

	time.Sleep(time.Second * 2)
	fmt.Println("writeOps", writeOps)
	fmt.Println("readOps", readOps)

}
