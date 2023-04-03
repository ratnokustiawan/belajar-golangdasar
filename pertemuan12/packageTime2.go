package main

import (
	"fmt"
	"time"
)

func main() {
	// waktu saat ini
	now := time.Now()
	fmt.Println("Waktu saat ini:", now)

	// menambahkan durasi ke waktu yang sudah ada
	duration, _ := time.ParseDuration("1h")
	oneHourLater := now.Add(duration)
	fmt.Println("Satu jam lagi:", oneHourLater)

	// mengurangi waktu dengan waktu yang lain
	newYear := time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)
	diff := newYear.Sub(now)
	fmt.Println("Sisa waktu menuju tahun baru:", diff)

	// memotong waktu menjadi waktu yang lebih kecil
	truncated := now.Truncate(24 * time.Hour)
	fmt.Println("Waktu setelah dipotong:", truncated)
}
