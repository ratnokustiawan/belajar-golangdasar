package main

import (
	"fmt"
)

func main() {
	hasil := sum(1, 2, 3, 4)
	fmt.Println(hasil)

	hasil = sum(100, 200, 300, 400, 500)
	fmt.Println(hasil)
}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
