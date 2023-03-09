package main

import "fmt"

func main() {
	// contoh penggunaan array dengan tipe data string
	var a [3]string
	a[0] = "Hello"
	a[1] = "World"
	a[2] = "!"
	fmt.Println("a:", a)
	fmt.Println("a[0]:", a[0])
	fmt.Println("a[1]:", a[1])
	fmt.Println("a[2]:", a[2])

	// contoh penggunaan array dengan tipe data integer
	var b [3]int
	b[0] = 1
	b[1] = 2
	b[2] = 3
	fmt.Println("b:", b)
	fmt.Println("b[0]:", b[0])
	fmt.Println("b[1]:", b[1])
	fmt.Println("b[2]:", b[2])

	// contoh penggunaan array dengan tipe data float
	var c [3]float64
	c[0] = 1.5
	c[1] = 2.5
	c[2] = 3.5
	fmt.Println("c:", c)
	fmt.Println("c[0]:", c[0])
	fmt.Println("c[1]:", c[1])
	fmt.Println("c[2]:", c[2])

	// contoh penggunaan array dengan tipe data boolean
	var d [3]bool
	d[0] = true
	d[1] = false
	d[2] = true
	fmt.Println("d:", d)
	fmt.Println("d[0]:", d[0])
	fmt.Println("d[1]:", d[1])
	fmt.Println("d[2]:", d[2])

	// contoh menghitung panjang array
	fmt.Println("Panjang array a adalah", len(a))
	fmt.Println("Panjang array b adalah", len(b))
	fmt.Println("Panjang array c adalah", len(c))
}
