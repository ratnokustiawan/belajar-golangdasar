package main

import "fmt"

func main() {
	var funcVar func(int) int

	increment := func(num int) int {
		return num + 1
	}

	funcVar = increment

	result := funcVar(10)

	fmt.Println("Result", result)
}
