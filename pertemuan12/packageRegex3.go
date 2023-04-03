package main

import (
	"fmt"
	"regexp"
)

func main() {
	r := regexp.MustCompile("a([a-z]+)e")
	submatch := r.FindStringSubmatch("apple")
	fmt.Println(submatch)
}
