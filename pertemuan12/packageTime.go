package main

import (
	"fmt"
	"time"
)

func main() {
	// waktu saat ini
	now := time.Now()
	fmt.Println("Waktu saat ini:", now)

	// mengubah string menjadi waktu
	layout := "2006-01-02"
	str := "2022-12-31"
	t, _ := time.Parse(layout, str)
	fmt.Println("Tanggal:", t)

	// mengubah waktu menjadi string
	layout = "02 Januari 2006"
	str2 := t.Format(layout)
	fmt.Println("Tanggal dalam string:", str2)
}
