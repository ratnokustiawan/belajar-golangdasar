package main

import "fmt"

func main() {
	defer fmt.Println("Ini dijalankan terakhir")
	fmt.Println("Ini dijalankan pertama")
}
