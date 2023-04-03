package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num int = 10
	value := reflect.ValueOf(num)
	fmt.Println("Nilai variabel num:", value.Int())
}
