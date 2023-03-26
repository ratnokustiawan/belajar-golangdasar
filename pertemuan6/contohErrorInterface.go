package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("Ini adalah contoh error")
	fmt.Println(err.Error())
}
