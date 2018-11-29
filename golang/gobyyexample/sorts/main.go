package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b", "d"}
	sort.Strings(strs)
	fmt.Println(strs)

	ints := []int{2, 3, 1, 4}
	sort.Ints(ints)
	fmt.Println(ints)
	fmt.Println(sort.IntsAreSorted(ints))
}
