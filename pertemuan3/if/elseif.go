package main

import "fmt"

func main() {
	var num = -1
	if num < 0 {
		fmt.Println("Negatif")
	} else if num > 0 {
		fmt.Println("Positif")
	} else {
		fmt.Print("nol")
	}
}
