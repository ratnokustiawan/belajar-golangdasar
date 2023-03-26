package main

import "fmt"

func main() {
	var a int = 10
	var ptr *int

	ptr = &a

	fmt.Printf("Nilai dari a: %d\n", a)
	fmt.Printf("Alamat dari a: %x\n", &a)
	fmt.Printf("Alamat dari ptr: %x\n", ptr)
	fmt.Printf("Nilai dari *ptr: %d\n", *ptr)
}
