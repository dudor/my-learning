package main

import (
	"fmt"
	"strings"
)

func main() {
	fruit := []string{"banana", "tea", "apple", "pear"}
	fmt.Println(Index(fruit, "tea"))
	fmt.Println(Include(fruit, "aaa"))
	fmt.Println(Filter(fruit, func(v string) bool {
		return strings.Contains(v, "e")
	}))

}
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func Include(vs []string, t string) bool {
	return Index(vs, t) > -1
}
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) == false {
			return false
		}
	}
	return true
}
func Filter(vs []string, f func(string) bool) []string {
	result := []string{}
	for _, v := range vs {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}
func Map(vs []string, f func(string) string) []string {
	result := make([]string, len(vs))
	for i, v := range vs {
		result[i] = f(v)
	}
	return result
}
