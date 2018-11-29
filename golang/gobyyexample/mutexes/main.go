package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync/atomic"
	"time"

	"sync"
)

func main() {
	state := make(map[int]int)
	mut := &sync.Mutex{}
	var ops int64 = 0
	total := 0

	for i := 0; i < 100; i++ {
		go func() {

			for {
				k := rand.Intn(5)
				mut.Lock()
				total += k
				mut.Unlock()
				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

	for i := 0; i < 100; i++ {
		go func() {
			for {
				k := rand.Intn(5)
				v := rand.Intn(100)
				mut.Lock()
				state[k] = v
				mut.Unlock()
				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second * 30)
	opsTotal := atomic.LoadInt64(&ops)
	fmt.Println("opstotal", opsTotal)
	mut.Lock()
	fmt.Println("state", state)
	fmt.Println("total", total)
	mut.Unlock()

}
