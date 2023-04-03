package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		name   string
		age    int
		isMale bool
	)

	flag.StringVar(&name, "name", "World", "a name to say hello to")
	flag.IntVar(&age, "age", 0, "an age")
	flag.BoolVar(&isMale, "male", false, "a gender")
	flag.Parse()

	fmt.Printf("Hello, my name is %s. I'm %d years old. I'm male: %t. \n", name, age, isMale)
}
