package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Mencocokkan sebuah pola pada sebuah string
	matched, _ := regexp.MatchString("a([a-z]+)e", "apple")
	fmt.Println(matched)

	// Mencocokkan sebuah pola pada sebuah string dan mengembalikan submatch
	r := regexp.MustCompile("a([a-z]+)e")
	submatch := r.FindStringSubmatch("apple")
	fmt.Println(submatch)
}
