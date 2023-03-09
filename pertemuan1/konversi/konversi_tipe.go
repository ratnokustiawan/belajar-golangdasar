package main

import "fmt"

func main() {
	var a string = "AHello, World!"
	var b byte = a[0]
	var c string = string(b)
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
}
