package main

import "fmt"

func main() {
	hari := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	slice := hari[1:]
	slice[0] = "Selasa Lagi"
	slice[1] = "Rabu Lagi"
	fmt.Println(hari)

	// append
	slice2 := append(slice, "Sabtu Lagi")
	slice2[0] = "Selasa Lagi Lagi"
	slice2[1] = "Rabu Lagi Lagi"
	fmt.Println(slice2)
	fmt.Println(hari)

	// // make
	slice3 := make([]string, 2, 5)
	slice3[0] = "Senin"
	slice3[1] = "Selasa"
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	// // copy
	copy(slice3, slice2)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
}
