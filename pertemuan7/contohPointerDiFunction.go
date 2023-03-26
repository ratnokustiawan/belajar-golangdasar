package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20

	fmt.Println("Sebelum swap, nilai a =", a, "dan nilai b =", b)

	swap(&a, &b)

	fmt.Println("Setelah swap, nilai a =", a, "dan nilai b =", b)
}

func swap(x *int, y *int) {
	var temp int

	temp = *x
	*x = *y
	*y = temp
}
