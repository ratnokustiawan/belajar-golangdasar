package main

import (
	"fmt"
	"sort"
)

func main() {
	// Mengurutkan slice integer
	a := []int{3, 2, 1}
	sort.Ints(a)
	fmt.Println(a) // Output: [1 2 3]

	// Mengecek apakah slice integer sudah terurut
	b := []int{1, 2, 3}
	fmt.Println(sort.IntsAreSorted(b)) // Output: true

	// Mengurutkan slice string
	c := []string{"c", "b", "a"}
	sort.Strings(c)
	fmt.Println(c) // Output: [a b c]

	// Mengecek apakah slice string sudah terurut
	d := []string{"a", "b", "c"}
	fmt.Println(sort.StringsAreSorted(d)) // Output: true
}
