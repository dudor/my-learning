package chanSync

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second * 2)
	fmt.Println("done..")
	done <- true
}

func Main() {
	done := make(chan bool, 1)
	go worker(done)
	<-done
}
