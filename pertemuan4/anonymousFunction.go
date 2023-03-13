package main

import "fmt"

func main() {
	func() {
		fmt.Println("Halo ini fungsi anonymous")
	}()

	jumlah := calculates(10, 5, func(x, y int) int {
		return x + y
	})

	kali := calculates(10, 5, func(x, y int) int {
		return x * y
	})

	fmt.Println("Jumlah", jumlah)
	fmt.Println("Kali", kali)
}

func calculates(x, y int, operation func(int, int) int) int {
	result := operation(x, y)
	return result
}
