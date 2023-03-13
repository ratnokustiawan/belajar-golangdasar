package main

import "fmt"

func main() {
	// do something
	hasil := tambah(10, 20)
	fmt.Println("10+20=", hasil)

	hasil = kurang(10, 20)
	fmt.Println("10+20=", hasil)
}

func tambah(a int, b int) int {
	return a + b
}

func kurang(a int, b int) int {
	return a - b
}
