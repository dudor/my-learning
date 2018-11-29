package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		t := strings.ToUpper(scan.Text())
		fmt.Println(t)
	}

	if err := scan.Err(); err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
}
