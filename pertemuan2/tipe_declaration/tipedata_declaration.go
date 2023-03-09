package main

import "fmt"

func main() {
	// named
	type Person string
	var nama Person = "Imron"
	fmt.Println(nama)

	//alias
	type PersonAlias = string
	var namaAlias PersonAlias = "Imron"
	fmt.Println(namaAlias)

	//struct

	type PersonStruct struct {
		nama string
		umur int
	}

	var orang PersonStruct = PersonStruct{nama: "Ratno", umur: 35}
	fmt.Println(orang)
}
