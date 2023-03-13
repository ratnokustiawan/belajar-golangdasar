package main

import "fmt"

func main() {
	hasil := factorial(5)
	fmt.Println("Hasil", hasil)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
