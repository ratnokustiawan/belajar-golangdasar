package main

import (
	"fmt"
	"regexp"
)

func main() {
	matched, _ := regexp.MatchString("a([a-z]+)e", "apple")
	fmt.Println(matched)
}
