package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Second(), now.Nanosecond(), now.Unix(), now.UnixNano())

}
