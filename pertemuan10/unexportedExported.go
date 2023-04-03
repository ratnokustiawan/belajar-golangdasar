package main

import "fmt"

// unexported variable
var unexportedVariable string = "Ini adalah variabel unexported"

// exported variable
var ExportedVariable string = "Ini adalah variabel exported"

func main() {
    fmt.Println(unexportedVariable)
    fmt.Println(ExportedVariable)
}
