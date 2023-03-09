package main

import "fmt"

func main() {
	arr := []string{"satu", "dua", "tiga", "empat"}
	for i := 0; i < len(arr); i++ {
		fmt.Println("data ke", i, "adalah", arr[i])
	}

}
