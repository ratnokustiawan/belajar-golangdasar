package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "ini adalah sebuah kalimat"
	count := 0

	for _, word := range strings.Split(str, " ") {
		count++
		fmt.Println(word)
	}

	fmt.Println(count)

}
