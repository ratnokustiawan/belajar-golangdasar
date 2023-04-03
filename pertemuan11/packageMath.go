package main

import (
	"fmt"
	"math"
)

func main() {
	// Menghitung nilai absolut dari -10
	fmt.Println("Nilai absolut dari -10 adalah", math.Abs(-10))

	// Membulatkan angka 3.14 ke atas
	fmt.Println("3.14 dibulatkan ke atas adalah", math.Ceil(3.14))

	// Menghitung nilai cosinus dari sudut 45 derajat
	fmt.Println("Cosinus dari sudut 45 derajat adalah", math.Cos(math.Pi/4))

	// Menghitung nilai eksponensial dari 2
	fmt.Println("Nilai eksponensial dari 2 adalah", math.Exp(2))

	// Membulatkan angka 3.14 ke bawah
	fmt.Println("3.14 dibulatkan ke bawah adalah", math.Floor(3.14))

	// Menghitung nilai logaritma dari 10 dengan basis 2
	fmt.Println("Nilai logaritma dari 10 dengan basis 2 adalah", math.Log2(10))

	// Mengembalikan nilai terbesar dari dua angka
	fmt.Println("Nilai terbesar dari 10 dan 20 adalah", math.Max(10, 20))

	// Mengembalikan nilai terkecil dari dua angka
	fmt.Println("Nilai terkecil dari 10 dan 20 adalah", math.Min(10, 20))

	// Menghitung nilai pangkat dari 2^3
	fmt.Println("2 pangkat 3 adalah", math.Pow(2, 3))

	// Menghitung nilai sinus dari sudut 30 derajat
	fmt.Println("Sinus dari sudut 30 derajat adalah", math.Sin(math.Pi/6))

	// Menghitung nilai akar kuadrat dari 16
	fmt.Println("Akar kuadrat dari 16 adalah", math.Sqrt(16))

	// Menghitung nilai tangen dari sudut 60 derajat
	fmt.Println("Tangen dari sudut 60 derajat adalah", math.Tan(math.Pi/3))
}
