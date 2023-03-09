package main

import "fmt"

func main() {
	str1 := "Golang"
	str2 := "golang"
	if str1 == str2 {
		fmt.Println("str1 sama dengan str2")
	} else {
		fmt.Println("str1 tidak sama dengan str2")
	}
}
