package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	r, e := json.Marshal(true)
	fmt.Println(r, e)
	r, e = json.Marshal("hello")
	fmt.Println(string(r), e)

	r2 := Response2{
		Page:   1,
		Fruits: []string{"apple", "banana"},
	}
	r, e = json.MarshalIndent(r2, "", "  ")
	fmt.Println(string(r), e)

	d := `"hello"`
	fmt.Printf("%T  %v", d, d)
}

type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}
