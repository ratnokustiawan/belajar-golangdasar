package main

import "fmt"

func main() {
	sum, difference := sumAndDifference(10, 20)
	fmt.Println("Sum", sum)
	fmt.Println("Diff", difference)
}

func sumAndDifference(a int, b int) (sum int, difference int) {
	sum = a + b
	difference = a - b
	return
}
