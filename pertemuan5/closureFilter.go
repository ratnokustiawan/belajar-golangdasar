package main

import "fmt"

func filter(numbers []int, f func(int) bool) []int {
	var result []int
	for _, v := range numbers {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	number := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	even := filter(number, func(x int) bool {
		return x%2 == 0
	})
	fmt.Println(even)
}
