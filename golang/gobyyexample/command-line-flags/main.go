package main

import (
	"flag"
	"fmt"
)

func main() {
	wortPtr := flag.String("word", "foo", "input a string")
	numPtr := flag.Int("num", 34, "an int")
	boolPtr := flag.Bool("fork", false, "a bool ")
	flag.Parse()
	fmt.Println(*wortPtr, *numPtr, *boolPtr)
	fmt.Println(flag.Args())
}
