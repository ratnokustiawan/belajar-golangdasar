package main

import "fmt"

func main() {
	a := 10
	b := 3

	fmt.Println("a & b = ", a&b)
	fmt.Println("a | b = ", a|b)
	fmt.Println("a ^ b = ", a^b)
	fmt.Println("^a = ", ^a)
	fmt.Println("a << 2 = ", a<<2)
	fmt.Println("a >> 2 = ", a>>2)
}
