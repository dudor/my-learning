package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("begin")
	f := createFile("test1.txt")
	defer closeFile(f)
	writeFile(f)
	fmt.Println("end")
}
func createFile(name string) *os.File {
	fmt.Println("creating")
	f, e := os.Create(name)
	if e != nil {
		panic(e)
	}
	return f
}
func writeFile(f *os.File) {
	fmt.Println("writing")
	f.WriteString("data")
}
func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}
