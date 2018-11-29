package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "hello,今天是周一。"
	r := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(r)
	r = base64.RawStdEncoding.EncodeToString([]byte(data))
	fmt.Println(r)
}
