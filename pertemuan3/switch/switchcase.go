package main

import "fmt"

func main() {
	var angka int = 8
	switch angka {
	case 1:
		fmt.Println("Hari ke-", angka, "adalah minggu")
	case 2:
		fmt.Println("Hari ke-", angka, "adalah senin")
	case 3:
		fmt.Println("Hari ke-", angka, "adalah selasa")
	case 4:
		fmt.Println("Hari ke-", angka, "adalah rabu")
	case 5:
		fmt.Println("Hari ke-", angka, "adalah kamis")
	case 6:
		fmt.Println("Hari ke-", angka, "adalah jumat")
	case 7:
		fmt.Println("Hari ke-", angka, "adalah sabtu")
	default:
		fmt.Println("angka yang anda masukkan tidak valid")
	}

}
