package main

import "fmt"

func main() {
	x := 5
	y := 10
	z := 15

	// and
	if x < y && y < z {
		fmt.Println("kondisi terpenuhi")
	} else {
		fmt.Println("Kondisi tidak terpenuhi")
	}

	// or
	if x > y || y < z {
		fmt.Println("Kondisi terpenuhi")
	} else {
		fmt.Println("Kondisi tidak terpenuhi")
	}

	//not

	if !(x > y) {
		fmt.Println("Kondisi terpenuhi")
	} else {
		fmt.Println("Kondisi tidak terpenuhi")
	}

}
