package main

import (
	"fmt"
	"os"
)

func main() {
	//os.Setenv("env1", "test1")
	for k, v := range os.Environ() {
		fmt.Println(k, v)
	}
}
