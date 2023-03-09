package main

import "fmt"

func main() {
	a := true
	b := false

	if a && b {
		fmt.Println("A dan B bernilai true")
	} else {
		fmt.Println("A dan B tidak bernilai true")
	}

	if a || b {
		fmt.Println("A atau B bernilai true")
	} else {
		fmt.Println("A atau B tidak bernilai true")
	}
}
