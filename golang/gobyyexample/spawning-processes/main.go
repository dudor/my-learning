package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd1 := exec.Command("go", "env")

	r, e := cmd1.Output()
	if e != nil {
		panic(e)
	}
	fmt.Println(string(r))
}
