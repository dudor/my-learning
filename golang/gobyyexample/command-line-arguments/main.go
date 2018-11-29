package main

import (
	"fmt"
	"os"
)

func main() {
	args_Pro := os.Args
	fmt.Println(args_Pro)
	fmt.Println(args_Pro[1:])
}
