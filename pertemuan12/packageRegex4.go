package main

import (
	"fmt"
	"regexp"
)

func main() {
	r := regexp.MustCompile("a([a-z]+)e")
	replaced := r.ReplaceAllString("apple", "orange")
	fmt.Println(replaced)
}
