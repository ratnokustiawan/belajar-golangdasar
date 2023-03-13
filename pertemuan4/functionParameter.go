package main

import "fmt"

func main() {
	x, y := 5, 10
	fmt.Println("Hasil jumlah", calculate(x, y, add))
	fmt.Println("Hasil kali", calculate(x, y, multiply))
}

func add(x int, y int) int {
	return x + y
}

func multiply(x int, y int) int {
	return x * y
}

func calculate(x int, y int, operation func(int, int) int) int {
	result := operation(x, y)
	return result
}
