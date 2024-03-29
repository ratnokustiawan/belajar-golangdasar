package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	a := adder()
	for i := 0; i < 5; i++ {
		fmt.Println(a(i))
	}
}
