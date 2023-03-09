package main

import "fmt"

func main() {
	a := 10
	b := 3

	sum := a + b
	sub := a - b
	mul := a * b
	div := a / b
	mod := a % b

	fmt.Println("Sum :", sum)
	fmt.Println("Sub :", sub)
	fmt.Println("Mul :", mul)
	fmt.Println("Div :", div)
	fmt.Println("Mod :", mod)
}
