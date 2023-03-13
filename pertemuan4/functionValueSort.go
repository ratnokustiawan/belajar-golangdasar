package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{2, 5, 4, 7, 9, 6}
	sort.Slice(numbers, func(i int, j int) bool {
		return numbers[i] < numbers[j]
	})
	fmt.Println(numbers)

}
