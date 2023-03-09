package main

import "fmt"

func main() {
	orang := map[string]string{
		"nama":   "Imron",
		"alamat": "Jl. Kebon Jeruk",
	}

	fmt.Println(orang)
	fmt.Println(orang["nama"])
	fmt.Println(orang["alamat"])

	delete(orang, "alamat")

	fmt.Println(orang)
}
