package main

import (
	"fmt"
	"pertemuan9/calculation"
)

func main() {
	var sisi float64 = 5
	luas := calculation.SquareArea(sisi)
	keliling := calculation.SquarePerimeter(sisi)

	fmt.Println("Luas persegi dengan sisi", sisi, "adalah", luas)
	fmt.Println("Keliling persegi dengan sisi", sisi, "adalah", keliling)
}
