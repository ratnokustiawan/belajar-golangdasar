package main

import (
	"container/ring"
	"fmt"
)

func main() {
	buffer := ring.New(3)

	for i := 1; i <= 5; i++ {
		buffer.Value = i
		buffer = buffer.Next()
	}

	buffer.Do(func(x interface{}) {
		fmt.Print(x, " ")
	})
}
