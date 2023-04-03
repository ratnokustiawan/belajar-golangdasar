package main

import (
	"fmt"
	"os"
)

func main() {
	hostname, _ := os.Hostname()
	fmt.Println("Hostname:", hostname)

	username := os.Getenv("USER")
	fmt.Println("Username:", username)

	wd, _ := os.Getwd()
	fmt.Println("Working directory:", wd)
}
