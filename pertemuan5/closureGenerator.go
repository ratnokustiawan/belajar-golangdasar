package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() func() int {
	rand.Seed(time.Now().UnixNano())
	return func() int {
		return rand.Intn(100)
	}
}

func main() {
	gen := generator()
	fmt.Println(gen())
	fmt.Println(gen())
}
