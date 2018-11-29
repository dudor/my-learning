package main

import (
	"fmt"
	"regexp"
)

func main() {
	m, e := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(m, e)
	r, e := regexp.Compile("p([a-z]+)ch")
	fmt.Println(r.MatchString("pacschch"))
}
