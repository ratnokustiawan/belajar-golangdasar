package main

import "fmt"

func main() {
	list := []interface{}{"Hello", 42, true}

	for _, val := range list {
		switch val.(type) {
		case string:
			fmt.Println(val.(string))
		case int:
			fmt.Println(val.(int))
		default:
			fmt.Println("Unknown type")
		}
	}
}
