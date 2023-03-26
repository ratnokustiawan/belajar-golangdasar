package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Terjadi error:", r)
		}
	}()
	panic("Terjadi kesalahan")
	fmt.Println("Selesai program")
}
