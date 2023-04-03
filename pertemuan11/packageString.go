package main

import (
	"fmt"
	"strings"
)

func main() {
	// Menggunakan function Contains
	fmt.Println(strings.Contains("Halo Dunia", "Dunia")) // Output: true
	fmt.Println(strings.Contains("Halo Dunia", "Hai"))   // Output: false

	// Menggunakan function Count
	fmt.Println(strings.Count("Dunia ini indah", "i")) // Output: 3

	// Menggunakan function Replace
	fmt.Println(strings.Replace("Halo Dunia", "Dunia", "Semua", 1)) // Output: Halo Semua
	fmt.Println(strings.Replace("Halo Dunia", "a", "i", -1))        // Output: Hilo Dunii

	// Menggunakan function ToUpper dan ToLower
	fmt.Println(strings.ToUpper("Halo Dunia")) // Output: HALO DUNIA
	fmt.Println(strings.ToLower("Halo Dunia")) // Output: halo dunia

	// Menggunakan function TrimSpace
	fmt.Println(strings.TrimSpace(" Halo Dunia ")) // Output: Halo Dunia
}
