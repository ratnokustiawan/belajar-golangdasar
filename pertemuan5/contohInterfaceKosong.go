package main

import "fmt"

func print(val interface{}) {
	fmt.Println(val)
}

func main() {
	print("Halo, dunia!") // Output: Halo, dunia!
	print(42)             // Output: 42
	print(3.14)           // Output: 3.14
}
