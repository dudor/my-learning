package main

import (
	"fmt"
	"sort"
)

func main()  {
	fruits := []string{"banana","apple","peach"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}

type ByLength []string

func (this ByLength)Len() int  {
	return len(this)
}
func (this ByLength)Swap(i,j int)  {
	this[i],this[j] = this[j],this[i]
}
func (this ByLength)Less(i,j int) bool {
	return len(this[i])<len(this[j])
}